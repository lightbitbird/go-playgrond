package command

import (
    "github.com/mkideal/cli"
)

type CliArgT struct {
    cli.Helper
    Port int `cli:"p,port" usage:"short and long format flags both are supported"`
    X bool `cli:"x" usage:"boolean type"`
    Y bool `cli:"y" usage:"boolean type, too"`
    Id uint8 `cli:"*id" usage:"this is a required flag, note the *"`
    DefaultDir string `cli:"dir" usage:"directory of developer" dft:"$HOME/dev"`
}
