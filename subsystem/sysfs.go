package subsystem

import (
	"github.com/mesanine/ginit/mount"
)

func Sysfs() Subsystem {
	data := "nodev,noexec,nosuid"
	return Subsystem{
		Mounts: []mount.MountArgs{
			mount.MountArgs{
				Source: "sysfs",
				Target: "/sys",
				FSType: "sysfs",
				Data:   data,
			},
			mount.MountArgs{
				Source: "securityfs",
				Target: "/sys/kernel/security",
				FSType: "securityfs",
				Data:   data,
			},
			mount.MountArgs{
				Source: "debugfs",
				Target: "/sys/kernel/debug",
				FSType: "debugfs",
				Data:   data,
			},
			mount.MountArgs{
				Source: "configfs",
				Target: "/sys/kernel/config",
				FSType: "configfs",
				Data:   data,
			},
			mount.MountArgs{
				Source: "fusectl",
				Target: "/sys/fs/fuse/connections",
				FSType: "fusectl",
				Data:   data,
			},
			mount.MountArgs{
				Source: "selinuxfs",
				Target: "/sys/fs/selinux",
				FSType: "selinuxfs",
				Data:   "nosuid,noexec",
			},
			mount.MountArgs{
				Source: "pstore",
				Target: "/sys/fs/pstore",
				FSType: "pstore",
				Data:   data,
			},
			mount.MountArgs{
				Source: "pstore",
				Target: "/sys/firmware/efi/efivars",
				FSType: "efivarfs",
				Data:   data,
			},
		},
	}
}
