package utils

import (
	"os"
	"syscall"
	"unsafe"

	"github.com/sirupsen/logrus"
)

var userdir string

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func GetWindowWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

func FileExists(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Size() > 0
}

func UserDir() string {
	if userdir == "" {
		var err error
		userdir, err = os.UserHomeDir()
		if err != nil {
			logrus.WithError(err).Fatalln("failed to get user home.")
		}
	}
	return userdir
}

func KubeConfig() string {
	return os.Getenv("KUBECONFIG")
}

func PrintSource(s string) {
	os.Stdout.WriteString(s)
}
