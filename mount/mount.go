/*
package mount contains helper functions for performing different
types of mounts in Linux.
*/
package mount

import (
	"golang.org/x/sys/unix"
)

// Before is a function called prior to
// performing a mount.
type Before func() error

// Option is a functional option to modify MountArgs.
type Option func(MountArgs) MountArgs

// Data changes the MountArgs Data prarameter.
func Data(data string) Option {
	return func(args MountArgs) MountArgs {
		return MountArgs{
			Source: args.Source,
			Target: args.Target,
			FSType: args.FSType,
			Flags:  args.Flags,
			// mount option flags, TODO: these should typed strongly
			Data:   data,
			Before: args.Before,
		}
	}
}

// MountArgs hold arguments for making
// unix.Mount syscalls.
type MountArgs struct {
	Source string
	Target string
	FSType string
	Flags  uintptr
	Data   string
	Before Before
}

// Mount performs the unix.Mount syscall
func Mount(args MountArgs, opts ...Option) error {
	for _, opt := range opts {
		args = opt(args)
	}
	if args.Before != nil {
		err := args.Before()
		if err != nil {
			return err
		}
	}
	return unix.Mount(args.Source, args.Target, args.FSType, args.Flags, args.Data)
}

// MustMount performs a unix.Mount syscall and panics on failure.
func MustMount(args MountArgs, opts ...Option) {
	err := Mount(args, opts...)
	if err != nil {
		panic(err)
	}
}
