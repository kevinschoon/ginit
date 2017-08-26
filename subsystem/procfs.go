package subsystem

import "github.com/mesanine/ginit/mount"

func ProcFS() Subsystem {
	return Subsystem{
		Mounts: []mount.MountArgs{mount.ProcFS()},
	}
}
