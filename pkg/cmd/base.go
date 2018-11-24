package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

const exitCodeNG = 1

var version = "unknown"
var revision string

// NewApp initializes *cli.App
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "svf"
	app.Usage = "A Shell Variables File Utility"
	app.Version = version
	app.Description = fmt.Sprintf("REVISION: %s", revision)
	app.Authors = []cli.Author{
		{
			Name:  "Shingo Kawamura",
			Email: "blp1526@gmail.com",
		},
	}
	app.Copyright = "(c) 2018 Shingo Kawamura"

	app.Commands = []cli.Command{
		keysCommand,
		valuesCommand,
	}

	return app
}
