package main

import (
    "fmt"
    "github.com/BurntSushi/toml"
    "os"
    "sync"
)

// CMD: go run toml_config.go {a config file name without an extension}
func main() {
    initEnv()
    url := Get().URL
    fmt.Println("url -> ", url)
}

type Config struct {
    // ここにTOMLと紐づけされる設定値を定義する。
    URL string `toml:"url"`
}

var instance *Config
var once sync.Once

func Get() *Config {
    return instance
}

func Init(e string) {
    once.Do(func() {
        env := e
        if e == "" {
            env = "development"
        }
        instance = &Config{}
        toml.DecodeFile("config/"+env+".toml", instance)
    })
}

func initEnv() {
    if len(os.Args) > 1 {
        Init(os.Args[1])
    } else {
        Init("")
    }
}
