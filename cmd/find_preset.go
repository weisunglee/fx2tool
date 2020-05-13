package main

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/Jockey66666/fx2preset-lite/pkg/filesystem"
	"github.com/Jockey66666/fx2preset-lite/pkg/pgjson"
)

// FindPreset : find preset by name
func FindPreset(rootPath, name string) {

	rootPath = filesystem.ExpandHome(rootPath)

	// open bank.json
	bankByte, err := filesystem.OpenFile(rootPath + "/bank.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var banks pgjson.BankList
	json.Unmarshal(bankByte, &banks)

	presetPath := ""
	for _, bank := range banks.LiveBanks {

		// open preset.json
		presetByte, _ := filesystem.OpenFile(rootPath + "/" + bank.BankFolder + "/preset.json")
		var presets pgjson.Preset
		json.Unmarshal(presetByte, &presets)

		for _, preset := range presets.LivePresets {
			if preset.PresetName == name {
				presetPath = rootPath + "/" + bank.BankFolder + "/" + preset.PresetUUID

				cmd := exec.Command("open", presetPath)
				cmd.Run()
			}
		}
	}
}
