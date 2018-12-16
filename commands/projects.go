package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

func newProjectsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "Get a list of all projects used",
		Long:  "Get a list of all projects used",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			frames, _ := chronolib.GetFrames(config)
			for _, project := range frames.Projects() {
				fmt.Println(project)
			}
		},
	}
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return cmd
}
