package main

import (
    "examples/go-playground/command"
    "github.com/mkideal/cli"
    "os"
)

// CMD example: go run cli_main.go --port{-p} 0000  -x=true -y=false -id=3 --dir xxx/xxx/*
func main() {
    os.Exit(cli.Run(new(command.CliArgT), func(ctx *cli.Context) error {
        argv := ctx.Argv().(*command.CliArgT)
        ctx.String("port=%d, x=%v, y=%v\nid=%d\ndev=%s", argv.Port, argv.X, argv.Y, argv.Id, argv.DefaultDir)
        return nil
    }))
}
