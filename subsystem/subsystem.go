/*
package subsystem contains helper modules for
initializing different Linux "sub systems" such
as /dev /proc, etc.
*/
package subsystem

import (
	"github.com/mesanine/ginit"
	"github.com/mesanine/ginit/mount"
)

type Subsystem struct {
	Nodes  []ginit.Node
	Links  []ginit.Link
	Mounts []mount.MountArgs
}

func Load(subs []Subsystem) error {
	for _, sub := range subs {
		for _, mnt := range sub.Mounts {
			err := mount.Mount(mnt)
			if err != nil {
				return err
			}
		}
		for _, node := range sub.Nodes {
			err := ginit.Mknod(node.Name, node.Mode, node.Type, node.Major, node.Minor)
			if err != nil {
				return err
			}
		}
		for _, link := range sub.Links {
			err := ginit.Mklink(link.Target, link.Name, link.Force)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
