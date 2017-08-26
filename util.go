package ginit

import (
	"golang.org/x/sys/unix"
	"os"
	"syscall"
)

func IsRoot() bool { return os.Getuid() == 0 }

// Exec does an execve with provided arguments
// it appends the executable to the front of
// the arguments and copies the existing environment.
func Exec(exe string, args ...string) error {
	newArgs := []string{exe}
	for _, arg := range args {
		newArgs = append(newArgs, arg)
	}
	return syscall.Exec(exe, newArgs, os.Environ())
}

// Check if a filesystem is memory based i.e. tempfs or ramfs
func IsMemFS(path string) (bool, error) {
	var stat unix.Statfs_t
	err := unix.Statfs(path, &stat)
	if err != nil {
		return false, err
	}
	if stat.Type == ramfsMagic || stat.Type == tmpfsMagic {
		return true, nil
	}
	return false, nil
}
