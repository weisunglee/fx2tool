package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Jockey66666/fx2preset-lite/pkg/filesystem"
)

const (
	settingPath = "/Documents/PositiveGrid/BIAS_FX2/settings.json"
)

// SwitchToProdction : switch to prodction
func SwitchToProdction() {
	userHome := filesystem.ExpandHome("~")
	os.Remove(userHome + settingPath)

	fmt.Println("Switch tonecloud to prodction.")
}

// SwitchToStaging : switch to staging
func SwitchToStaging() {

	stagingSetting := map[string]string{
		"Dev Settings.PaymentMode": "Staging",
	}

	userHome := filesystem.ExpandHome("~")
	json, _ := json.MarshalIndent(stagingSetting, "", "    ")
	ioutil.WriteFile(userHome+settingPath, json, 0644)

	fmt.Println("Switch tonecloud to staging.")
}
