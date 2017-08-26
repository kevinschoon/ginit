package ginit

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Controller represents a cgroup system controller
// see http://man7.org/linux/man-pages/man7/cgroups.7.html
type Controller struct {
	Name       string
	Hierarchy  int
	NumCgroups int
	Enabled    bool
}

// ReadControllers returns Cgroup controllers listed at /proc/cgroups
// Modified from https://github.com/opencontainers/runc/blob/master/libcontainer/cgroups/utils.go
func ReadControllers() ([]Controller, error) {
	f, err := os.Open("/proc/cgroups")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	controllers := []Controller{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		text := s.Text()
		if text[0] != '#' {
			parts := strings.Fields(text)
			if len(parts) >= 4 && parts[3] != "0" {
				cont := Controller{
					Name: parts[0],
				}
				n, err := strconv.ParseInt(parts[1], 0, 64)
				if err == nil {
					cont.Hierarchy = int(n)
				}
				n, err = strconv.ParseInt(parts[2], 0, 64)
				if err == nil {
					cont.NumCgroups = int(n)
				}
				e, err := strconv.ParseBool(parts[3])
				if err == nil {
					cont.Enabled = e
				}
				controllers = append(controllers, cont)
			}
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return controllers, nil
}
