package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_41"
	app.Usage = "cli_41 sample"

	// before
	app.Before = func(c *cli.Context) error {
		fmt.Println("-- Before --")
		return nil
	}

	// command action
	app.Commands = []cli.Command{
		// `go run cli_04/main.go sample パラメータ`でactionするコマンド
		{
			Name:    "sample",
			Aliases: []string{"s"},
			Usage:   "sample task",
			Action: func(c *cli.Context) error {
				fmt.Println("-- Sample Action --")
				fmt.Println("sample task: ", c.Args().First())
				return nil
			},
		},
	}

	// action
	app.Action = func(c *cli.Context) error {
		fmt.Println("-- Nomal Action --")
		return nil
	}

	// after
	app.After = func(c *cli.Context) error {
		fmt.Println("-- After --")
		return nil
	}
	app.Run(os.Args)
}
