package ginit

import "golang.org/x/sys/unix"

// Node represents a Linux "node" or "file" in a way that is simplier
// than the unix.Stat_t type.
type Node struct {
	Name  string
	Mode  uint32
	Type  uint32
	Major int64
	Minor int64
}

// Mkdev is used to build the value of linux devices (in /dev/) which specifies major
// and minor number of the newly created device special file.
// Linux device nodes are a bit weird due to backwards compat with 16 bit device nodes.
// They are, from low to high: the lower 8 bits of the minor, then 12 bits of the major,
// then the top 12 bits of the minor.
// From https://github.com/moby/moby/blob/master/pkg/system/mknod.go
func Mkdev(major int64, minor int64) uint32 {
	return uint32(((minor & 0xfff00) << 12) | ((major & 0xfff) << 8) | (minor & 0xff))
}

// Mknod creates a block or character special file with
// arguments similar to gnu coreutils.
// Example:
// Mknod("/dev/console", 0600, unix.S_IFCHR, int64(5), int64(1))
// $ mknod -m 600 /dev/console c 5 1
func Mknod(name string, mode, ntype uint32, major, minor int64) error {
	return unix.Mknod(name, mode|ntype, int(Mkdev(major, minor)))
}
