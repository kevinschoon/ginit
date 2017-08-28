package ginit

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
)

// Bind returns arguments for perofmring a Bind mount.
func Bind(path string, readOnly bool) MountArgs {
	var flags uintptr
	if readOnly {
		flags = unix.MS_BIND | unix.MS_RDONLY
	} else {
		flags = unix.MS_BIND
	}
	return MountArgs{
		Source: path,
		Target: path,
		Flags:  flags,
	}
}

// Overlay returns options for performing
// an OverlayFS mount.
func Overlay(lower, target string) MountArgs {
	upper := filepath.Join(filepath.Dir(lower), "upper")
	work := filepath.Join(filepath.Dir(lower), "work")
	return MountArgs{
		Before: func() (err error) {
			err = os.MkdirAll(upper, 0755)
			if err != nil {
				return err
			}
			err = os.MkdirAll(work, 0755)
			if err != nil {
				return err
			}
			return nil
		},
		Source: "overlay",
		Target: target,
		FSType: "overlay",
		Flags:  0,
		Data:   fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lower, upper, work),
	}
}

// TmpFS return options for performing a tmpfs mount
// at the given path percentage must be between 0
// and 100 or we will panic. If it is zero we do
// not specify any flags.
func TmpFS(path string, percentage int) MountArgs {
	if percentage < 0 || percentage > 100 {
		panic("invalid tempfs percentage")
	}
	var data string
	if percentage > 0 {
		data = fmt.Sprintf("%d", percentage)
	}
	return MountArgs{
		Before: func() error {
			return os.MkdirAll(path, 0755)
		},
		Source: "rootfs", // TODO: Unsure if this has significance with tempfs
		Target: path,
		FSType: "tmpfs",
		Data:   data,
	}
}
