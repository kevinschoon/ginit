package ginit

import (
	"golang.org/x/sys/unix"
	"os"
)

// Is checks if a file at the path has the given mode.
// If the file does not exist it returns false.
func Is(path string, mode os.FileMode) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.Mode() == mode, nil
}

// IsMust checks if a file at the path has a given mode.
// If it encounters an error the function will panic.
func IsMust(path string, mode os.FileMode) bool {
	v, err := Is(path, mode)
	if err != nil {
		panic(err)
	}
	return v
}

// Stat returns linux file or device status
func Stat(path string) (unix.Stat_t, error) {
	info, err := os.Stat(path)
	if err != nil {
		return unix.Stat_t{}, err
	}
	stat := *info.Sys().(*unix.Stat_t)
	return stat, nil
}

// StatMust returns file or device status
// or panics if an error is encountered.
func StatMust(path string) unix.Stat_t {
	stat, err := Stat(path)
	if err != nil {
		panic(err)
	}
	return stat
}
