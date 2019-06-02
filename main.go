package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"sdkman-cli/sdkmanCli"
)

func main() {
	env := &sdkmanCli.Env{"~\\sdkman", "C:\\Users\\palin\\.sdkman\\candidates", []string{"java", "scala"}, false, "https://api.sdkman.io/2/candidates", "1"}
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name: "list",
			Aliases: []string{
				"l", "ls",
			},
			Action: func(c *cli.Context) error {
				sdkmanCli.List(c.Args().First(), env)
				return nil
			},
		}, {
			Name:    "current",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) error {
				sdkmanCli.Current(env, c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
