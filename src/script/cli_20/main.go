package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_02"
	app.Usage = "cli_02 sample"
	app.Version = "1.2.1"

	// before
	app.Before = func(c *cli.Context) error {
		fmt.Println("-- Before --")
		return nil
	}

	// action
	app.Action = func(c *cli.Context) error {
		fmt.Println("-- Action --")

		fmt.Printf("c.NArg()        : %+v\n", c.NArg())
		fmt.Printf("c.Args()        : %+v\n", c.Args())
		fmt.Printf("c.Args().Get(0) : %+v\n", c.Args().Get(0))
		fmt.Printf("c.Args()[0]     : %+v\n", c.Args()[0])
		fmt.Printf("c.String()      : %+v\n", c.String("game"))

		// ここでHelp表示
		//cli.ShowAppHelp(c)

		// ここでversionを表示
		cli.ShowVersion(c)

		return nil
	}

	// after
	app.After = func(c *cli.Context) error {
		fmt.Println("-- After --")
		return nil
	}

	app.HideHelp = false

	app.Run(os.Args)
}
