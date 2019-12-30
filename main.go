package main

import (
	"github.com/palindrom615/sdkman-cli/command"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Name:  "sdkman",
		Usage: "manage various versions of SDKs",
		Commands: []*cli.Command{
			{
				Name: "list",
				Aliases: []string{
					"l", "ls",
				},
				Usage: "[candidate]",
				Action: func(c *cli.Context) error {
					return command.List(c.Args().First())
				},
			}, {
				Name:    "current",
				Aliases: []string{"c"},
				Usage:   "[candidate]",
				Action: func(c *cli.Context) error {
					return command.Current(c.Args().First())
				},
			}, {
				Name:  "update",
				Usage: "",
				Action: func(c *cli.Context) error {
					return command.Update()
				},
			}, {
				Name:    "install",
				Usage:   "<candidate> [version]",
				Aliases: []string{"i"},
				Action: func(c *cli.Context) error {
					return command.Install(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
				},
			}, {
				Name:  "use",
				Usage: "<candidate> <version>",
				Action: func(c *cli.Context) error {
					return command.Use(c.Args().Get(0), c.Args().Get(1))
				},
			}, {
				Name:  "export",
				Usage: "[shell]",
				Action: func(c *cli.Context) error {
					return command.Export(c.Args().Get(0))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
