package subsystem

import (
	"github.com/mesanine/ginit"
)

func ProcFS() ginit.Subsystem {
	return ginit.Subsystem{
		Mounts: []ginit.MountArgs{ginit.ProcFS()},
	}
}
