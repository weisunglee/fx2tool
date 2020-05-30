package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/Jockey66666/fx2tool/pkg/filesystem"
)

// RestorePreset : factory reset
func RestorePreset() {
	userHome := filesystem.ExpandHome("~")

	if runtime.GOOS == "windows" {
		removeAllPresets(userHome)
	} else {
		copyPresetFromBundle(userHome)
	}

	removeAmpList(userHome)
	removeFxList(userHome)
	removeSetting()
}

func copyPresetFromBundle(userHome string) {
	const src = `/Applications/BIAS FX 2.app/Contents/Resources/HTML/LiveDefaultPresets/`
	dst := userHome + "/Documents/PositiveGrid/BIAS_FX2/GlobalPresets/"

	os.RemoveAll(dst)
	err := filesystem.CopyFolder(src, dst)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Restore preset to")
		fmt.Println(dst)
		fmt.Println()
	}
}

func removeSetting() {
	userSetting := filesystem.GetSettingsPath() + "/PositiveGrid/BIAS_FX2.settings"

	err := os.Remove(userSetting)
	if err == nil {
		fmt.Println(userSetting, "is removed.")
	}
}

func removeAmpList(userHome string) {
	amplist := userHome + "/Documents/PositiveGrid/BIAS_FX2/AmpList.json"

	err := os.Remove(amplist)
	if err == nil {
		fmt.Println(amplist, "is removed.")
	}
}

func removeFxList(userHome string) {
	fxlist := userHome + "/Documents/PositiveGrid/BIAS_FX2/FxList.json"

	err := os.Remove(fxlist)
	if err == nil {
		fmt.Println(fxlist, "is removed.")
	}
}

func removeAllPresets(userHome string) {
	dst := userHome + "/Documents/PositiveGrid/BIAS_FX2/GlobalPresets/"
	os.RemoveAll(dst)
}
