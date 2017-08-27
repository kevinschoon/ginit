package ginit

type Subsystem struct {
	Nodes  []Node
	Links  []Link
	Mounts []MountArgs
}

func Load(subs []Subsystem) error {
	for _, sub := range subs {
		for _, mnt := range sub.Mounts {
			err := Mount(mnt)
			if err != nil {
				return err
			}
		}
		for _, node := range sub.Nodes {
			err := Mknod(node.Name, node.Mode, node.Type, node.Major, node.Minor)
			if err != nil {
				return err
			}
		}
		for _, link := range sub.Links {
			err := Mklink(link.Target, link.Name, link.Force)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
