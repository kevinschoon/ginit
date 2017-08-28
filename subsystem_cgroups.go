package ginit

import (
	"fmt"
)

func Cgroups(controllers []Controller) Subsystem {
	mounts := []MountArgs{
		MountArgs{
			Target: "/sys/fs/cgroup",
			Source: "cgroup_root",
			FSType: "tempfs",
			Data:   "nodev,noexec,nosuid,mode=755,size=10m",
		},
	}
	for _, controller := range controllers {
		if controller.Enabled {
			mounts = append(mounts, MountArgs{
				Target: fmt.Sprintf("/sys/fs/cgroup/%s", controller.Name),
				Source: controller.Name,
				FSType: "cgroup",
				Data:   fmt.Sprintf("nodev,noexec,nosuid,%s", controller.Name),
			})
		}
	}
	return Subsystem{
		Mounts: mounts,
	}
}
