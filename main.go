package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	sdkman_cli "sdkman-cli/sdkman-cli"
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
				sdkman_cli.ListSdk(c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
