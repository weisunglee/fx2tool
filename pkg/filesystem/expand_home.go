// +build darwin linux

package filesystem

import (
	"os/user"
	"path/filepath"
)

// ExpandHome : Expand home dir
func ExpandHome(path string) string {
	if len(path) == 0 || path[0] != '~' {
		return path
	}

	usr, err := user.Current()
	if err != nil {
		return ""
	}
	return filepath.Join(usr.HomeDir, path[1:])
}
