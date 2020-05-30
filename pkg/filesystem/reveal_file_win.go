// +build windows

package filesystem

import (
	"os/exec"
	"strings"
)

// RevealFile : reveal in explorer
func RevealFile(presetPath string) {
	presetPath = strings.ReplaceAll(presetPath, "/", "\\")

	cmd := exec.Command("explorer", presetPath)
	cmd.Run()
}
