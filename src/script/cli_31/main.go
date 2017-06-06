package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_31"
	app.Usage = "cli_31 sample"
	app.Version = "1.2.1"

	// flags
	app.Flags = []cli.Flag{
		// StringFlag
		cli.StringFlag{
			Name:  "name, n",
			Value: "tarou",
			Usage: "your name",
		},
		// BoolFlag
		cli.BoolFlag{
			Name:  "gay, g",
			Usage: "are you gay boy?",
		},
	}

	// action
	app.Action = func(c *cli.Context) error {
		fmt.Println("-- Action --")
		fmt.Printf("c.GlobalFlagNames() : %+v\n", c.GlobalFlagNames())
		fmt.Printf("c.String(\"name\")    : %+v\n", c.String("name"))
		fmt.Printf("c.String(\"g\")       : %+v\n", c.Bool("g"))

		return nil
	}

	app.Run(os.Args)
}
