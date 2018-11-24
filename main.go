package main

import (
	"os"

	"github.com/blp1526/go-shell-variables-file/pkg/cmd"
)

func main() {
	app := cmd.NewApp()
	app.Run(os.Args)
}
