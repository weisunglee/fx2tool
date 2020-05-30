package cmd

import (
	"fmt"
	"os"

	"github.com/Jockey66666/fx2tool/pkg/filesystem"
)

// Logout : log out from tonecloud
func Logout() {

	userHome := filesystem.ExpandHome("~")
	settings := []string{
		".license",
		".jwt_token",
		".pgbiasfx2",
		".pgbiasfx2_pa_info_setting",
		".pgbiasfx2_user_setting",
	}

	for _, s := range settings {
		file := userHome + "/Documents/PositiveGrid/BIAS_FX2/" + s

		err := os.Remove(file)
		if err == nil {
			fmt.Println(file, "is removed")
		}
	}
}
