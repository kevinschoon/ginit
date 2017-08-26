package ginit

import (
	"os"
)

// Link describes a symlink
type Link struct {
	Target string
	Name   string
	Force  bool
}

// Mklink creates a symlink link optionally deleteing existing files
func Mklink(target, name string, force bool) error {
	if force {
		err := os.Remove(name)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}
	}
	return os.Symlink(target, name)
}
