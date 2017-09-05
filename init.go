package ginit

import (
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"time"
)

// InitArgs configure the runtime of Init
type InitArgs struct{}

// Init launches a signal handler loop.
// Resources:
// https://github.com/mirror/busybox/tree/master/init
// https://github.com/torvalds/linux/blob/master/kernel/reboot.c
func Init(args InitArgs) error {
	var n int
	sigCh := make(chan os.Signal, 10)
	signal.Reset()
	signal.Notify(sigCh)
loop:
	for {
		sig := <-sigCh
		switch sig {
		// Halt
		case unix.SIGUSR1:
			n = unix.LINUX_REBOOT_CMD_HALT
			break loop
		// Poweroff
		case unix.SIGUSR2:
			n = unix.LINUX_REBOOT_CMD_POWER_OFF
			break loop
		// Reboot
		case unix.SIGTERM:
			n = unix.LINUX_REBOOT_CMD_RESTART
			break loop
		}
	}
	// SIGTERM all processes except pid 1
	err := unix.Kill(-1, unix.SIGTERM)
	if err != nil {
		return err
	}
	time.Sleep(1)
	// SIGKILL all processes except pid 1
	err = unix.Kill(-1, unix.SIGKILL)
	if err != nil {
		return err
	}
	time.Sleep(1)
	unix.Sync()
	return unix.Reboot(n)
}
