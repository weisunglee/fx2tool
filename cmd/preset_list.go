package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Jockey66666/fx2tool/pkg/filesystem"
	"github.com/Jockey66666/fx2tool/pkg/pgjson"

	"github.com/fatih/color"
)

var magenta = color.New(color.FgHiMagenta).SprintFunc()
var red = color.New(color.FgHiRed).SprintFunc()
var cyan = color.New(color.FgHiCyan).SprintFunc()
var yellow = color.New(color.FgHiYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()

// ListAllPresets : list all presets
func ListAllPresets(root string) {

	root = filesystem.ExpandHome(root)

	// open bank.json
	bankByte, err := filesystem.OpenFile(root + "/bank.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var banks pgjson.BankList
	json.Unmarshal(bankByte, &banks)

	fmt.Println(root)
	totalPresets := 0
	for i, bank := range banks.LiveBanks {
		bankPrefix := "├──"

		if i == len(banks.LiveBanks)-1 {
			bankPrefix = "└──"
		}

		fmt.Printf("%s %s - %s\n", bankPrefix, magenta(bank.BankFolder), cyan(bank.BankName))

		// open preset.json
		presetByte, err := filesystem.OpenFile(root + "/" + bank.BankFolder + "/preset.json")

		if err != nil {
			os.Exit(0)
		}

		var presets pgjson.Preset
		json.Unmarshal(presetByte, &presets)

		parent := "│"
		for _, preset := range presets.LivePresets {

			if i == len(banks.LiveBanks)-1 {
				parent = " "
			}

			presetPrefix := parent + "   ├──"

			fmt.Printf("%s %s - %s\n", presetPrefix, red(preset.PresetUUID), yellow(preset.PresetName))
		}

		if i < len(banks.LiveBanks)-1 {
			fmt.Print("│")
		} else {
			fmt.Print(" ")
		}

		fmt.Printf("   └── %d preset(s) in total.\n", len(presets.LivePresets))
		totalPresets += len(presets.LivePresets)
	}

	fmt.Println("")
	fmt.Printf("%d preset(s) in %d bank(s).\n", totalPresets, len(banks.LiveBanks))
}
