// +build windows

package filesystem

import (
	"path/filepath"
	"syscall"

	"github.com/lxn/win"
)

// ExpandHome : Expand home dir
func ExpandHome(path string) string {
	if len(path) == 0 || path[0] != '~' {
		return path
	}

	return filepath.Join(userHomeDir(), path[1:])
}

func userHomeDir() string {
	buf := make([]uint16, win.MAX_PATH)
	win.SHGetSpecialFolderPath(win.HWND(0), &buf[0], win.CSIDL_PROFILE, false)
	path := syscall.UTF16ToString(buf)
	return path
}
