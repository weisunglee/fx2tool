// +build darwin linux

package filesystem

import "os/exec"

// RevealFile : reveal in finder
func RevealFile(presetPath string) {
	cmd := exec.Command("open", presetPath)
	cmd.Run()
}
