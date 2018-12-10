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
			stateStorage := chronolib.GetStateStorage(config)
			state, err := stateStorage.Get()
			if err != nil {
				commandError = err
			} else {
				if state.Project == "" {
					fmt.Println(chronolib.FormatNoProjectMessage())
				} else {
					fmt.Println(chronolib.RenderStatusFormatString(state, format))
				}
			}
		},
	}

	newStatus.Flags().StringVarP(&format, "format", "f", defaultFormat, "go template string to format output")
	return newStatus
}
