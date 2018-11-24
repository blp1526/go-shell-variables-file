package cmd

import (
	"fmt"

	"github.com/blp1526/go-shell-variables-file/pkg/svf"
	"github.com/urfave/cli"
)

var keysCommand = cli.Command{
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
			value, _ := s.GetValue(key)
			fmt.Println(value)
		}

		return nil
	},
}
