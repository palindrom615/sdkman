package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"sdkman-cli/sdkmanCli"
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
				sdkmanCli.List(c.Args().First())
				return nil
			},
		}, {
			Name:    "current",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) error {
				sdkmanCli.Current(c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
