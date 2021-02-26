package goeject

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"unsafe"
)

func SgioSyscall(f *os.File, i *SgIoHdr) error {
	return ioctl(f.Fd(), SG_IO, uintptr(unsafe.Pointer(i)))
}

func ioctl(fd, cmd, ptr uintptr) error {
	_, _, err := unix.Syscall(unix.SYS_IOCTL, fd, cmd, ptr)
	if err != 0 {
		return err
	}
	return nil
}

func NewFile(deviceFilename string) (*os.File, error) {

	f, err := os.OpenFile(deviceFilename, os.O_RDONLY|unix.O_NONBLOCK, 0660)
	if err != nil {
		return nil, err
	}
	var version uint32
	if (ioctl(f.Fd(), SG_GET_VERSION_NUM, uintptr(unsafe.Pointer(&version))) != nil) || (version < 30000) {
		return nil, fmt.Errorf("device does not appear to be an sg device")
	}
	return f, nil

}