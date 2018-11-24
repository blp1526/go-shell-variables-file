package cmd

import (
	"fmt"

	"github.com/blp1526/go-shell-variables-file/pkg/svf"
	"github.com/urfave/cli"
)

const exitCodeNG = 1

var version = "unknown"
var revision string

// NewApp initializes *cli.App
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "svf"
	app.Version = version
	app.Description = fmt.Sprintf("REVISION: %s", revision)
	app.Authors = []cli.Author{
		{
			Name:  "Shingo Kawamura",
			Email: "blp1526@gmail.com",
		},
	}
	app.Copyright = "(c) 2018 Shingo Kawamura"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Usage: "set a shell variables file path",
		},

		cli.BoolFlag{
			Name:  "list",
			Usage: "show keys list",
		},

		cli.StringFlag{
			Name:  "output",
			Usage: "output `key`'s value",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "keys",
			Usage:     "show keys",
			ArgsUsage: "path",
			Action: func(c *cli.Context) (err error) {
				path := c.Args().First()
				s, err := svf.ReadFile(path)
				if err != nil {
					return cli.NewExitError(err, exitCodeNG)
				}

				for _, key := range s.Keys() {
					fmt.Println(key)
				}

				return nil
			},
		},

		{
			Name:      "values",
			Usage:     "show values",
			ArgsUsage: "path",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key",
					Usage: "show a value by a key",
				},
			},
			Action: func(c *cli.Context) (err error) {
				path := c.Args().First()
				s, err := svf.ReadFile(path)
				if err != nil {
					return cli.NewExitError(err, exitCodeNG)
				}

				keys := s.Keys()
				key := c.String("key")
				if key != "" {
					keys = []string{key}
				}

				err = s.IsValidKeys(keys)
				if err != nil {
					return cli.NewExitError(err, exitCodeNG)
				}

				for _, key := range keys {
					value, _ := s.Value(key)
					fmt.Println(value)
				}

				return nil
			},
		},
	}

	return app
}
