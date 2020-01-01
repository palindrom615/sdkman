package main

import (
	"github.com/palindrom615/sdkman"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path"
)

func main() {
	home, _ := os.UserHomeDir()
	app := &cli.App{
		Name:  "sdkman",
		Usage: "manage various versions of SDKs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "directory",
				Aliases: []string{"d"},
				Usage:   "directory to save SDKs",
				Value:   path.Join(home, ".sdkman"),
			},
			&cli.StringFlag{
				Name:    "registry",
				Aliases: []string{"reg"},
				Usage:   "sdkman server url",
				Value:   "https://api.sdkman.io/2",
			},
			&cli.BoolFlag{
				Name:  "insecure",
				Usage: "ignore ssl certificate error",
				Value: false,
			},
		},
		Commands: []*cli.Command{
			{
				Name: "list",
				Aliases: []string{
					"l", "ls",
				},
				Usage:  "[candidate]",
				Action: sdkman.List,
			}, {
				Name:    "current",
				Aliases: []string{"c"},
				Usage:   "[candidate]",
				Action:  sdkman.Current,
			}, {
				Name:   "update",
				Usage:  "",
				Action: sdkman.Update,
			}, {
				Name:    "install",
				Usage:   "<candidate> [version]",
				Aliases: []string{"i"},
				Action:  sdkman.Install,
			}, {
				Name:   "use",
				Usage:  "<candidate> <version>",
				Action: sdkman.Use,
			}, {
				Name:   "export",
				Usage:  "[shell]",
				Action: sdkman.Export,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
