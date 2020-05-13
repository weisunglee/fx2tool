package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "fx2preset-lite"
	app.Usage = "command line utility for BIAS FX2 preset data"
	app.Version = "1.0.0"

	const rootPath = "~/Documents/PositiveGrid/BIAS_FX2/GlobalPresets"

	app.Commands = []*cli.Command{
		{
			Name:    "restore",
			Aliases: []string{"r"},
			Usage:   "Restore documents and settings",
			Action: func(c *cli.Context) error {
				RestorePreset()
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all presets",
			Action: func(c *cli.Context) error {
				ListAllPresets(rootPath)
				return nil
			},
		},
		{
			Name:    "find",
			Aliases: []string{"f"},
			Usage:   "Find preset name",
			Action: func(c *cli.Context) error {
				FindPreset(rootPath, c.Args().First())
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
