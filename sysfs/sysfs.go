package sysfs

import (
	"github.com/mesanine/ginit"
)

func Subsystem() ginit.Subsystem {
	data := "nodev,noexec,nosuid"
	return ginit.Subsystem{
		Mounts: []ginit.MountArgs{
			ginit.MountArgs{
				Source: "sysfs",
				Target: "/sys",
				FSType: "sysfs",
				Data:   data,
			},
			ginit.MountArgs{
				Source: "securityfs",
				Target: "/sys/kernel/security",
				FSType: "securityfs",
				Data:   data,
			},
			ginit.MountArgs{
				Source: "debugfs",
				Target: "/sys/kernel/debug",
				FSType: "debugfs",
				Data:   data,
			},
			ginit.MountArgs{
				Source: "configfs",
				Target: "/sys/kernel/config",
				FSType: "configfs",
				Data:   data,
			},
			ginit.MountArgs{
				Source: "fusectl",
				Target: "/sys/fs/fuse/connections",
				FSType: "fusectl",
				Data:   data,
			},
			ginit.MountArgs{
				Source: "selinuxfs",
				Target: "/sys/fs/selinux",
				FSType: "selinuxfs",
				Data:   "nosuid,noexec",
			},
			ginit.MountArgs{
				Source: "pstore",
				Target: "/sys/fs/pstore",
				FSType: "pstore",
				Data:   data,
			},
			ginit.MountArgs{
				Source: "pstore",
				Target: "/sys/firmware/efi/efivars",
				FSType: "efivarfs",
				Data:   data,
			},
		},
	}
}
