package ginit

import (
	"fmt"
	//"golang.org/x/sys/unix"
	"net"
	"strings"
)

// Hostname generates a hostname based on the
// prefix and the first valid MAC address from sysfs.
func Hostname(prefix string) (string, error) {
	sfs := KeyFS{Base: "/sys"}
	keys, err := sfs.Find("/class/net", "address", true)
	if err != nil {
		return "", err
	}
	for _, key := range keys {
		raw, err := sfs.Read(key)
		if err != nil {
			return "", err
		}
		value := strings.Replace(string(raw), "\n", "", -1)
		// Skip loopback
		if value == "00:00:00:00:00:00" {
			continue
		}
		mac, err := net.ParseMAC(string(value))
		if err == nil {
			return fmt.Sprintf("%s-%s", prefix, mac.String()), nil
		}
	}
	return "", fmt.Errorf("unable to generate hostname")
}
