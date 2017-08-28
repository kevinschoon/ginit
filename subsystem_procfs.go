package ginit

func ProcFS() Subsystem {
	return Subsystem{
		Mounts: []MountArgs{
			MountArgs{
				Source: "proc",
				Target: "/proc",
				FSType: "proc",
				Flags:  0,
				Data:   "nodev,nosuid,noexec,relatime",
			},
		},
	}
}
