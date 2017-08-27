package ginit

// Node represents a Linux "node" or "file" in a way that is simplier
// than the unix.Stat_t type.
type Node struct {
	Name  string
	Mode  uint32
	Type  uint32
	Major uint32
	Minor uint32
}
