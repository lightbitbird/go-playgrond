package main

import "examples/go-playground/config"

// CMD: go run viper_main.go -c {a config file name without full path} -n xxxx -d
func main() {
    config.Execute()
}
