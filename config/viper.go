package config

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "os"
)

type Config struct {
    ApplicationName string
    Debug bool
}

var configFile string
var config Config

var rootCmd = &cobra.Command {
    Use: "app",
    Run: func(c *cobra.Command, args []string) {
        fmt.Printf("configFile: %s\nconfig: %#v\n", configFile, config)
    },
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config/config.toml", "config file name")
    rootCmd.PersistentFlags().StringVarP(&config.ApplicationName, "name", "n", "", "application name")
    rootCmd.PersistentFlags().BoolVarP(&config.Debug, "debug", "d", false, "debug mode")
    viper.BindPFlag("ApplicationName", rootCmd.PersistentFlags().Lookup("name"))
    viper.BindPFlag("Debug", rootCmd.PersistentFlags().Lookup("debug"))

    cobra.OnInitialize(func() {
        viper.SetConfigFile(configFile)
        // overwrite Env values
        viper.AutomaticEnv()

        if err := viper.ReadInConfig(); err != nil {
            fmt.Println("config file read error")
            fmt.Println(err)
            os.Exit(1)
        }

        if err := viper.Unmarshal(&config); err != nil {
            fmt.Println("config file Unmarshal error")
            fmt.Println(err)
            os.Exit(1)
        }
    })
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
