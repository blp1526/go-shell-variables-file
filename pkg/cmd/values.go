package cmd

import (
	"fmt"

	"github.com/blp1526/go-shell-variables-file/pkg/svf"
	"github.com/urfave/cli"
)

var valuesCommand = cli.Command{
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
}
