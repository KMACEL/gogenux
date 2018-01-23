package operations

import (
	"log"
	"syscall"
)

// Reboot is
func Reboot() {
	errs := syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART2)
	log.Println("Error :", errs.Error())
}
