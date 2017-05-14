package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_03"
	app.Usage = "cli_03 sample"
	app.Version = "1.2.1"

	os.Setenv("SAMPLE_ENV", "sample env dayo")

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name:  "meridian, m",
			Value: "AM",
			Usage: "meridian for the greeting",
		},
		cli.StringFlag{
			Name:  "time, t",
			Value: "07:00",
			// ``で囲むとhelp時のPlaceholderとしても使える
			// https://github.com/urfave/cli#placeholder-values
			Usage: "`your time` for the greeting",
		},
		cli.StringFlag{
			Name:  "aaa, a",
			Value: "sample",
			// default値をValueからではなくEnvから取る
			EnvVar: "SAMPLE_ENV",
		},
	}

	// action
	app.Action = func(c *cli.Context) error {
		fmt.Println("-- Action --")

		fmt.Printf("c.GlobalFlagNames() : %+v\n", c.GlobalFlagNames())
		fmt.Printf("c.String(\"lang\")    : %+v\n", c.String("lang"))
		fmt.Printf("c.String(\"m\")       : %+v\n", c.String("m"))
		fmt.Printf("c.String(\"time\")    : %+v\n", c.String("time"))
		fmt.Printf("c.String(\"a\")       : %+v\n", c.String("a"))

		return nil
	}

	app.Run(os.Args)
}
