package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_20"
	app.Usage = "cli_20 sample"
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

		// Help表示
		//cli.ShowAppHelp(c)

		// versionを表示
		cli.ShowVersion(c)

		return nil
	}

	// after
	app.After = func(c *cli.Context) error {
		fmt.Println("-- After --")
		return nil
	}

	app.HideHelp = true

	app.Run(os.Args)
}
