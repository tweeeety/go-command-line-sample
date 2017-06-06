package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_43"
	app.Usage = "cli_43 sample"

	// command action
	app.Commands = []cli.Command{
		// addコマンドに対してFlagsを設定
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				fmt.Println("json      : ", c.String("j"))
				fmt.Println("exec      : ", c.String("e"))
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "json, j",
					Value: "add.json",
				},
				cli.BoolFlag{
					Name: "exec, e",
				},
			},
		},
		// completeコマンドに対してFlagsを設定
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				fmt.Println("csv           : ", c.String("c"))
				fmt.Println("exec          : ", c.String("e"))
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "csv, c",
					Value: "add.json",
				},
				cli.BoolFlag{
					Name: "exec, e",
				},
			},
		},
	}
	app.Run(os.Args)
}
