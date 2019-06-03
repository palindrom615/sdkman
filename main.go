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
			Action: func(c *cli.Context) error {
				command.List(c.Args().First())
				return nil
			},
		}, {
			Name:    "current",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) error {
				command.Current(c.Args().First())
				return nil
			},
		}, {
			Name: "update",
			Action: func(c *cli.Context) error {
				command.Update()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
