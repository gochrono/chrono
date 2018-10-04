package commands

import (
    "fmt"
    "strings"
    "github.com/spf13/cobra"
    "github.com/jordanknott/chrono/chronolib"
    humanize "github.com/dustin/go-humanize"
)
func newStatusCmd() *cobra.Command {
    return &cobra.Command{
        Use: "status",
        Short: "Get status of current frame",
        Long: "Get the status of the current frame",
        Run: func(cmd *cobra.Command, args []string) {
            statePath := chronolib.GetAppFilePath("state", "")
            frame := chronolib.LoadState(statePath)
            if ( frame.Project == "" ) {
                fmt.Println("No project started")
            } else {
                started := frame.StartedAt
                fmt.Println("Project " + frame.Project + " started " + humanize.Time(started) + " [" + strings.Join(frame.Tags, ", ") + "]")
            }

        },
    }
}
