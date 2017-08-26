package subsystem

import (
	"github.com/mesanine/ginit"
	"github.com/mesanine/ginit/mount"
	"golang.org/x/sys/unix"
)

func DevFS() Subsystem {
	return Subsystem{
		Mounts: []mount.MountArgs{
			mount.MountArgs{
				Source: "dev",
				Target: "/dev",
				FSType: "devtmpfs",
				// TODO: What are nr_inodes?
				// TODO: What should be configurable here?
				Data: "nodev,noexec,relatime,size=10m,nr_inodes=248418,mode=755",
			},
			mount.MountArgs{
				Source: "mqueue",
				Target: "/dev/mqueue",
				FSType: "mqueue",
				Data:   "noexec,nosuid,nodev",
			},
			mount.MountArgs{
				Source: "shm",
				Target: "/dev/shm",
				FSType: "tmpfs",
				Data:   "noexec,nosuid,nodev,mode=1777",
			},
			mount.MountArgs{
				Source: "devpts",
				Target: "/dev/pts",
				FSType: "devpts",
				Data:   "noexec,nosuid,gid=5,mode=0620",
			},
		},
		Nodes: []ginit.Node{
			ginit.Node{
				Name:  "/dev/console",
				Mode:  0600,
				Type:  unix.S_IFCHR,
				Major: 5,
				Minor: 1,
			},
			ginit.Node{
				Name:  "/dev/console",
				Mode:  0620,
				Type:  unix.S_IFCHR,
				Major: 4,
				Minor: 1,
			},
			ginit.Node{
				Name:  "/dev/console",
				Mode:  0666,
				Type:  unix.S_IFCHR,
				Major: 5,
				Minor: 0,
			},
			ginit.Node{
				Name:  "/dev/console",
				Mode:  0666,
				Type:  unix.S_IFCHR,
				Major: 1,
				Minor: 3,
			},
			ginit.Node{
				Name:  "/dev/console",
				Mode:  0660,
				Type:  unix.S_IFCHR,
				Major: 1,
				Minor: 11,
			},
		},
		Links: []ginit.Link{
			ginit.Link{
				Name:   "/proc/self/fd",
				Target: "/dev/stdin",
				Force:  true,
			},
			ginit.Link{
				Name:   "/proc/self/fd/0",
				Target: "/dev/stdout",
				Force:  true,
			},
			ginit.Link{
				Name:   "/proc/self/fd/1",
				Target: "/dev/stdout",
				Force:  true,
			},
			ginit.Link{
				Name:   "/proc/self/fd/2",
				Target: "/dev/stderr",
				Force:  true,
			},
		},
	}
}
