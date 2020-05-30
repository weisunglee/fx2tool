// +build windows

package filesystem

import (
	"syscall"

	"github.com/lxn/win"
)

// GetSettingsPath : get settings path
func GetSettingsPath() string {
	buf := make([]uint16, win.MAX_PATH)
	win.SHGetSpecialFolderPath(win.HWND(0), &buf[0], win.CSIDL_APPDATA, false)
	path := syscall.UTF16ToString(buf)
	return path
}
