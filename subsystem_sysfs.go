package ginit

func Sysfs() Subsystem {
	data := "nodev,noexec,nosuid"
	return Subsystem{
		Mounts: []MountArgs{
			MountArgs{
				Source: "sysfs",
				Target: "/sys",
				FSType: "sysfs",
				Data:   data,
			},
			MountArgs{
				Source: "securityfs",
				Target: "/sys/kernel/security",
				FSType: "securityfs",
				Data:   data,
			},
			MountArgs{
				Source: "debugfs",
				Target: "/sys/kernel/debug",
				FSType: "debugfs",
				Data:   data,
			},
			MountArgs{
				Source: "configfs",
				Target: "/sys/kernel/config",
				FSType: "configfs",
				Data:   data,
			},
			MountArgs{
				Source: "fusectl",
				Target: "/sys/fs/fuse/connections",
				FSType: "fusectl",
				Data:   data,
			},
			MountArgs{
				Source: "selinuxfs",
				Target: "/sys/fs/selinux",
				FSType: "selinuxfs",
				Data:   "nosuid,noexec",
			},
			MountArgs{
				Source: "pstore",
				Target: "/sys/fs/pstore",
				FSType: "pstore",
				Data:   data,
			},
			MountArgs{
				Source: "pstore",
				Target: "/sys/firmware/efi/efivars",
				FSType: "efivarfs",
				Data:   data,
			},
		},
	}
}
