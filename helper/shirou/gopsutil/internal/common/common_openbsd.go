//go:build openbsd
// +build openbsd

package common

import (
	"os"
	"os/exec"
	"strings"
	"unsafe"

	"golang.org/x/sys/unix"
)

func DoSysctrl(mib string) ([]string, error) {
	sysctl, err := exec.LookPath("/sbin/sysctl")
	if err != nil {
		return []string{}, err
	}
	cmd := exec.Command(sysctl, "-n", mib)
	cmd.Env = getSysctrlEnv(os.Environ())
	out, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}
	v := strings.Replace(string(out), "{ ", "", 1)
	v = strings.Replace(string(v), " }", "", 1)
	values := strings.Fields(string(v))

	return values, nil
}

func CallSyscall(mib []int32) ([]byte, uint64, error) {
	mibptr := unsafe.Pointer(&mib[0])
	miblen := uint64(len(mib))

	// get required buffer size
	length := uint64(0)
	_, _, err := unix.Syscall6(
		unix.SYS___SYSCTL,
		uintptr(mibptr),
		uintptr(miblen),
		0,
		uintptr(unsafe.Pointer(&length)),
		0,
		0)
	if err != 0 {
		var b []byte
		return b, length, err
	}
	if length == 0 {
		var b []byte
		return b, length, err
	}
	// get proc info itself
	buf := make([]byte, length)
	_, _, err = unix.Syscall6(
		unix.SYS___SYSCTL,
		uintptr(mibptr),
		uintptr(miblen),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&length)),
		0,
		0)
	if err != 0 {
		return buf, length, err
	}

	return buf, length, nil
}
