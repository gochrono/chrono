package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var banner = `   _____ _
  / ____| |
 | |    | |__  _ __ ___  _ __   ___
 | |    | '_ \| '__/ _ \| '_ \ / _ \ 
 | |____| | | | | | (_) | | | | (_) |
  \_____|_| |_|_|  \___/|_| |_|\___/`

func newVersionCmd() *cobra.Command {
    return &cobra.Command{
        Use: "version",
        Short: "Prints the version then exits",
        Long: "Prints the version then exits",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println(banner)
            fmt.Printf("\nversion: %s\ncommit: %s\nbuilt: %s\n", version, commit, date)
        },
    }
}
