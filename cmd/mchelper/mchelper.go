package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "mchelper"
	app.Version = "0.0.0"
	app.Author = "Jared Allard <jaredallard@outlook.com>"
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "run",
			Description: "execute a command on the minecraft server",
			Usage:       "run <command>",
			Action:      runCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
