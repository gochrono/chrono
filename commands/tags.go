package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

func newTagsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tags",
		Short: "Get a list of all tags used",
		Long:  "Get a list of all tags used",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			frames, _ := chronolib.GetFrames(config)
			tags := frames.Tags()
			for _, tag := range tags {
				fmt.Println(tag)
			}
		},
	}
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return cmd
}
