package ginit

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// KeyFS implements a simple structure for
// interacting with "key" based filesystems
// like procfs and sysfs.
type KeyFS struct {
	Base string
}

// Find traverses the path below the Base path
// and optionally follows one level of symbolic
// link. For example, to find the mac addresses
// of all network devices you could do:
// Find("/class/net", "address", true)
// And it will return an array of resolved devices
// at /sys/devices/pci...
func (k KeyFS) Find(dir, name string, follow bool) ([]string, error) {
	var keys []string
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}
		if (info.Mode()&os.ModeSymlink) == os.ModeSymlink && follow {
			err := os.Chdir(filepath.Dir(path))
			if err != nil {
				return err
			}
			path, err := os.Readlink(path)
			if err != nil {
				return err
			}
			path, err = filepath.Abs(path)
			if err != nil {
				return err
			}
			found, err := k.Find(strings.Replace(path, k.Base, "", -1), name, false)
			for _, key := range found {
				keys = append(keys, key)
			}
		}
		if _, n := filepath.Split(path); n == name {
			keys = append(keys, strings.Replace(path, k.Base, "", -1))
		}
		return nil
	}
	err := filepath.Walk(filepath.Join(k.Base, dir), walkFn)
	return keys, err
}

// Read reads from the given key
func (k KeyFS) Read(key string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(k.Base, key))
}

// Write writes to the given key
func (k KeyFS) Write(key string, value []byte) error {
	stat, err := os.Stat(filepath.Join(k.Base, key))
	if err != nil {
		return err
	}
	fd, err := os.OpenFile(filepath.Join(k.Base, key), os.O_WRONLY, stat.Mode())
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = fd.Write(value)
	return err
}
