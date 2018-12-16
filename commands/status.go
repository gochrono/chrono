package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

const defaultFormat = "Project {{ .Project | magenta }}{{ if .Tags }} [{{ joinTags .Tags | blue }}]{{ end }} started {{ humanize .StartedAt | green }}."

var format string

func newStatusCmd() *cobra.Command {
	newStatus := &cobra.Command{
		Use:   "status",
		Short: "Get status of current frame",
		Long:  "Get the status of the current frame",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			state, err := chronolib.GetState(config)
			if err != nil {
				panic(err)
			}
			if state.IsEmpty() {
				fmt.Println(chronolib.FormatNoProjectMessage())
			} else {
				fmt.Println(chronolib.RenderCurrentFrameStatus(state.Get(), format))
			}
		},
	}

	newStatus.Flags().StringVarP(&format, "format", "f", defaultFormat, "go template string to format output")
	newStatus.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return newStatus
}
