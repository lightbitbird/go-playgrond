package main

import (
    "examples/go-playground/command"
)

// CMD: go run command_main.go sub --name{-n} xxx -d
func main() {
    command.Execute()
}
