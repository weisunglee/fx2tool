package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Name = "fx2-tool"
	app.Usage = "command line utility for BIAS FX2 preset data"
	app.Version = "1.0.5"

	const rootPath = "~/Documents/PositiveGrid/BIAS_FX2/GlobalPresets"

	app.Commands = []*cli.Command{
		{
			Name:    "restore",
			Aliases: []string{"r"},
			Usage:   "Restore preset from bundle and remove BIAS_FX2.settings",
			Action: func(c *cli.Context) error {
				RestorePreset()
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Show all presets",
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
		{
			Name:    "logout",
			Aliases: []string{"lo"},
			Usage:   "Log out from tonecloud",
			Action: func(c *cli.Context) error {
				Logout()
				return nil
			},
		},
		{
			Name:    "env",
			Aliases: []string{"e"},
			Usage:   "Switch tonecloud to production/staging",
			Subcommands: []*cli.Command{
				{
					Name:    "prod",
					Aliases: []string{"p"},
					Usage:   "prodction environment",
					Action: func(c *cli.Context) error {
						SwitchToProdction()
						return nil
					},
				},
				{
					Name:    "stage",
					Aliases: []string{"s"},
					Usage:   "staging environment",
					Action: func(c *cli.Context) error {
						SwitchToStaging()
						return nil
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
