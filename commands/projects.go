package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

func newProjectsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "projects",
		Short: "Get a list of all projects used",
		Long:  "Get a list of all projects used",
		Run: func(cmd *cobra.Command, args []string) {
            configDir := chronolib.GetCorrectConfigDirectory("")
            config := chronolib.GetConfig(configDir)
			frameStorage := chronolib.GetFrameStorage(config)
			projects, err := frameStorage.Projects()
			if err != nil {
				commandError = err
				return
			}

			for _, project := range projects {
				fmt.Println(project)
			}
		},
	}
}