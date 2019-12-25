package main

import (
	"github.com/palindrom615/sdkman-cli/command"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "list",
				Aliases: []string{
					"l", "ls",
				},
				Action: func(c *cli.Context) error {
					return command.List(c.Args().First())
				},
			}, {
				Name:    "current",
				Aliases: []string{"c"},
				Action: func(c *cli.Context) error {
					return command.Current(c.Args().First())
				},
			}, {
				Name: "update",
				Action: func(c *cli.Context) error {
					return command.Update()
				},
			}, {
				Name:    "install",
				Aliases: []string{"i"},
				Action: func(c *cli.Context) error {
					return command.Install(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
				},
			}, {
				Name: "use",
				Action: func(c *cli.Context) error {
					return command.Use(c.Args().Get(0), c.Args().Get(1))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
