package main

import (
	"github.com/palindrom615/sdk/command"
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
				Action: command.List,
			}, {
				Name:    "current",
				Aliases: []string{"c"},
				Usage:   "[candidate]",
				Action:  command.Current,
			}, {
				Name:   "update",
				Usage:  "",
				Action: command.Update,
			}, {
				Name:    "install",
				Usage:   "<candidate> [version]",
				Aliases: []string{"i"},
				Action:  command.Install,
			}, {
				Name:   "use",
				Usage:  "<candidate> <version>",
				Action: command.Use,
			}, {
				Name:   "export",
				Usage:  "[shell]",
				Action: command.Export,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
