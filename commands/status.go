package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/jordanknott/chrono/chronolib"
)
const defaultFormat = "Project {{ .Project | magenta }}{{ if .Tags }} [{{ joinTags .Tags | blue }}]{{ end }} started {{ humanize .StartedAt | green }}."

var format string

func newStatusCmd() *cobra.Command {
    newStatus := &cobra.Command{
        Use: "status",
        Short: "Get status of current frame",
        Long: "Get the status of the current frame",
        Run: func(cmd *cobra.Command, args []string) {
            statePath := chronolib.GetAppFilePath("state", "")
            frame := chronolib.LoadState(statePath)
            if ( frame.Project == "" ) {
                fmt.Println("No project started")
            } else {
                fmt.Println(chronolib.RenderStatusFormatString(*frame, format))
            }
        },
    }

    newStatus.Flags().StringVarP(&format, "format", "f", defaultFormat, "go template string to format output")
    return newStatus
}
