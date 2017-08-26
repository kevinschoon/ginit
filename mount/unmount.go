package mount

import "golang.org/x/sys/unix"

// Unmount unmounts a file system at the specified path
func Unmount(path string) error {
	return unix.Unmount(path, 0)
}
