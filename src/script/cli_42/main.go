package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_42"
	app.Usage = "cli_42 sample"

	// command action
	app.Commands = []cli.Command{
		{
			Name:    "sample",
			Aliases: []string{"s"},
			Usage:   "sample task",
			Action: func(c *cli.Context) error {
				fmt.Println("-- Sample Action --")

				fmt.Printf("c.Command.FullName()        : %+v\n", c.Command.FullName())
				fmt.Printf("c.Command.HasName(\"sample\") : %+v\n", c.Command.HasName("sample"))
				fmt.Printf("c.Command.Names()           : %+v\n", c.Command.Names())
				fmt.Printf("c.Command.VisibleFlags()    : %+v\n", c.Command.VisibleFlags())

				return nil
			},
		},
	}

	app.Run(os.Args)
}
