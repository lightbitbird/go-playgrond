package command

import (
    "fmt"
    "github.com/spf13/cobra"
    "log"
    "os"
)

var debug bool

var rootCmd = &cobra.Command{
    Use: "app",
    Run: func(c *cobra.Command, args []string) {
        fmt.Println("debug:", debug)
    },
}

func init() {
    rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug enable flag")
    subCmd.PersistentFlags().String("name", "Tom", "sub command string flag test")
    rootCmd.AddCommand(subCmd)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

var subCmd = &cobra.Command{
    Use: "sub",
    Run: func(c *cobra.Command, args []string) {
        if name, err := c.PersistentFlags().GetString("name"); err == nil {
            if debug {
                log.Println("From log name: ", name)
            } else {
                fmt.Println("name: ", name)
            }
        }
    },
}
