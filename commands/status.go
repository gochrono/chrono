package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/jordanknott/chrono/chronolib"
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
                fmt.Println(chronolib.FormatStatusFrameMessage(*frame))
            }
        },
    }
}
