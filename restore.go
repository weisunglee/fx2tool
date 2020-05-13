package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jockey66666/fx2preset-lite/pkg/filesystem"
)

// RestorePreset : factory reset
func RestorePreset() {
	userHome := filesystem.ExpandHome("~")
	copyPresetFromBundle(userHome)
	removeListSettings(userHome)
	removeSetting(userHome)
}

func copyPresetFromBundle(userHome string) {
	const src = `/Applications/BIAS FX 2.app/Contents/Resources/HTML/LiveDefaultPresets/`
	dst := userHome + "/Documents/PositiveGrid/BIAS_FX2/GlobalPresets/"

	err := filesystem.CopyFolder(src, dst)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Restore preset to")
		fmt.Println(dst)
		fmt.Println()
	}
}

func removeSetting(userHome string) {
	userSetting := userHome + "/Library/Application Support/PositiveGrid/BIAS_FX2.settings"
	os.Remove(userSetting)
	fmt.Println(userSetting, "is removed.")
}

func removeListSettings(userHome string) {
	amplist := userHome + "/Documents/PositiveGrid/BIAS_FX2/AmpList.json"
	os.Remove(amplist)
	fmt.Println(amplist, "is removed.")

	fxlist := userHome + "/Documents/PositiveGrid/BIAS_FX2/FxList.json"
	os.Remove(fxlist)
	fmt.Println(fxlist, "is removed.")
}
