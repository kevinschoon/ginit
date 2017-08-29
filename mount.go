package ginit

import (
	"golang.org/x/sys/unix"
)

// Option is a functional option to modify MountArgs.
type MountOption func(MountArgs) MountArgs

// MountArgs hold arguments for making
// unix.Mount syscalls.
type MountArgs struct {
	Source string
	Target string
	FSType string
	Flags  uintptr
	// mount -o options
	// TODO: Make strongly typed.
	Data   string
	Before func() error
}

// Mount performs the unix.Mount syscall
func Mount(args MountArgs, opts ...MountOption) error {
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
func MustMount(args MountArgs, opts ...MountOption) {
	err := Mount(args, opts...)
	if err != nil {
		panic(err)
	}
}
