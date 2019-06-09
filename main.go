package main

import (
	"log"
	"os"
	"sdkman-cli/command"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name: "list",
			Aliases: []string{
				"l", "ls",
			},
			Action: func(c *cli.Context) {
				command.List(c.Args().First())
			},
		}, {
			Name:    "current",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) {
				command.Current(c.Args().First())
			},
		}, {
			Name: "update",
			Action: func(c *cli.Context) {
				command.Update()
			},
		}, {
			Name:    "install",
			Aliases: []string{"i"},
			Action: func(c *cli.Context) {
				command.Install(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
			},
		}, {
			Name: "use",
			Action: func(c *cli.Context) {
				command.Use(c.Args().Get(0), c.Args().Get(1))
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
