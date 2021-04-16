package main

import (
	"os"

	"github.com/mkideal/cli"
)

type argT struct {
	Name string `cli:"name" usage:"tell me your name"`
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("Hello, %s!\n", argv.Name)
		return nil
	}))
}
