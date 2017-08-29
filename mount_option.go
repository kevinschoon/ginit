package ginit

import (
	"os"
)

// Mkdir creates the target directory
// if it is missing before performing
// a mount operation.
func Mkdir(mode os.FileMode) MountOption {
	return func(args MountArgs) MountArgs {
		return MountArgs{
			Source: args.Source,
			Target: args.Target,
			FSType: args.FSType,
			Data:   args.Data,
			Before: func() error {
				if args.Before != nil {
					err := args.Before()
					if err != nil {
						return err
					}
				}
				return os.MkdirAll(args.Target, mode)
			},
		}
	}
}

// Data changes the MountArgs Data prarameter.
func Data(data string) MountOption {
	return func(args MountArgs) MountArgs {
		return MountArgs{
			Source: args.Source,
			Target: args.Target,
			FSType: args.FSType,
			Flags:  args.Flags,
			Data:   data,
			Before: args.Before,
		}
	}
}
