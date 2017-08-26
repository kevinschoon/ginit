/*
Most of the code here is just copied over from Linuxkit's init package,
made slightly more readable, and for consumption as a library.
https://github.com/linuxkit/linuxkit/blob/master/pkg/init/cmd/init/init.go
*/
package ginit

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

const (
	ramfsMagic    = 0x858458f6
	tmpfsMagic    = 0x01021994
	switchBufSize = 32768
)

type SwitchOptions struct {
	BufSize int
	RootDev uint64
	NewRoot string
}

func NewSwitchOptions(path string) (*SwitchOptions, error) {
	// find the device of the root filesystem so we can avoid changing filesystem
	info, err := os.Stat("/")
	if err != nil {
		return nil, err
	}
	return &SwitchOptions{
		NewRoot: path,
		BufSize: switchBufSize,
		RootDev: info.Sys().(*syscall.Stat_t).Dev,
	}, nil
}

func copyTree(opts SwitchOptions) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// skip non directories
		if !info.Mode().IsDir() {
			return nil
		}
		dest := filepath.Join(opts.NewRoot, path)
		// create the directory
		if path == "/" {
			// the mountpoint already exists but may have wrong mode, metadata
			if err := os.Chmod(opts.NewRoot, info.Mode()); err != nil {
				return err
			}
		} else {
			if err := os.Mkdir(dest, info.Mode()); err != nil {
				return err
			}
		}
		if err := CopyFileInfo(info, dest); err != nil {
			return err
		}
		// skip recurse into other filesystems
		stat := info.Sys().(*syscall.Stat_t)
		if opts.RootDev != stat.Dev {
			return filepath.SkipDir
		}
		return nil
	}
}

func copyFiles(opts SwitchOptions) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// skip other filesystems
		stat := info.Sys().(*syscall.Stat_t)
		if opts.RootDev != stat.Dev && info.Mode().IsDir() {
			return filepath.SkipDir
		}
		dest := filepath.Join(opts.NewRoot, path)
		buf := make([]byte, opts.BufSize)
		switch {
		case info.Mode().IsDir():
			// already done the directories
			return nil
		case info.Mode().IsRegular():
			// TODO support hard links (currently not handled well in initramfs)
			new, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, info.Mode())
			if err != nil {
				return err
			}
			old, err := os.Open(path)
			if err != nil {
				return err
			}
			if _, err := io.CopyBuffer(new, old, buf); err != nil {
				return err
			}
			if err := old.Close(); err != nil {
				return err
			}
			if err := new.Close(); err != nil {
				return err
			}
			// it is ok if we do not remove all files now
			err = os.Remove(path)
			if err != nil {
				// TODO: Remove this. Just curious to see what cannot be deleted.
				fmt.Printf("error removing file: %s %s", path, err.Error())
			}
		case (info.Mode() & os.ModeSymlink) == os.ModeSymlink:
			link, err := os.Readlink(path)
			if err != nil {
				return err
			}
			if err := os.Symlink(link, dest); err != nil {
				return err
			}
		case (info.Mode() & os.ModeDevice) == os.ModeDevice:
			if err := unix.Mknod(dest, uint32(info.Mode()), int(stat.Rdev)); err != nil {
				return err
			}
		case (info.Mode() & os.ModeNamedPipe) == os.ModeNamedPipe:
			// TODO support named pipes, although no real use case
			return errors.New("Unsupported named pipe on rootfs")
		case (info.Mode() & os.ModeSocket) == os.ModeSocket:
			// TODO support sockets, although no real use case
			return errors.New("Unsupported socket on rootfs")
		default:
			return errors.New("Unknown file type")
		}
		if err := CopyFileInfo(info, dest); err != nil {
			return err
		}
		// TODO copy extended attributes if needed
		return nil
	}
}

func removeFiles(opts SwitchOptions) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ignore root itself
		if path == "/" {
			return nil
		}
		switch {
		case info.Mode().IsDir():
			// skip other filesystems (ie newRoot)
			stat := info.Sys().(*syscall.Stat_t)
			if opts.RootDev != stat.Dev {
				return filepath.SkipDir
			}
			// do our best to delete
			err = os.RemoveAll(path)
			if err != nil {
				// TODO: Remove this. Just curious to see what cannot be deleted.
				fmt.Printf("error removing file: %s %s", path, err.Error())
			}
			return filepath.SkipDir
		default:
			err = os.RemoveAll(path)
			if err != nil {
				// TODO: Remove this. Just curious to see what cannot be deleted.
				fmt.Printf("error removing file: %s %s", path, err.Error())
			}
			return nil
		}
	}
}

// CopyFileInfo takes an os.FileInfo and applies it
// to a file at the given path.
func CopyFileInfo(info os.FileInfo, path string) error {
	// would rather use fd than path but Go makes this very difficult at present
	stat := info.Sys().(*syscall.Stat_t)
	if err := unix.Lchown(path, int(stat.Uid), int(stat.Gid)); err != nil {
		return err
	}
	timespec := []unix.Timespec{unix.Timespec(stat.Atim), unix.Timespec(stat.Mtim)}
	if err := unix.UtimesNanoAt(unix.AT_FDCWD, path, timespec, unix.AT_SYMLINK_NOFOLLOW); err != nil {
		return err
	}
	// after chown suid bits may be dropped; re-set on non symlink files
	if info.Mode()&os.ModeSymlink == 0 {
		if err := os.Chmod(path, info.Mode()); err != nil {
			return err
		}
	}
	return nil
}

// SwitchRoot performs a memory efficient
// file copy of the root file system onto
// a new mount point and then pivots to
// the new location. Opts.NewRoot must
// already be mounted for this to work.
func SwitchRoot(opts SwitchOptions) error {
	// Copy the directory tree of the current
	// root path into the new mount point
	err := filepath.Walk("/", copyTree(opts))
	if err != nil {
		return err
	}
	// Copy all the files of the root path
	// into the new mount point
	err = filepath.Walk("/", copyFiles(opts))
	if err != nil {
		return err
	}
	// chdir to the new root directory
	if err := os.Chdir(opts.NewRoot); err != nil {
		return err
	}
	// mount --move cwd (/mnt) to /
	if err := unix.Mount(".", "/", "", unix.MS_MOVE, ""); err != nil {
		return err
	}
	// chroot to .
	if err := unix.Chroot("."); err != nil {
		return err
	}
	// chdir to "/" to fix up . and ..
	if err := os.Chdir("/"); err != nil {
		return err
	}
	return nil
}
