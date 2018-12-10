package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

var startNote string
var startAt string
var startEnded string

func newStartCmd() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start [project name] [tags]",
		Short: "Start a new frame",
		Long: `Starts a new frame with the given project name and tags.
		Tags must start with a plus (+) to be considered valid. Any spaces will be converted to
		dashes (-). It will also be lowercased. To stop the timer for the current frame, use
		the 'stop' command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			stateStorage := chronolib.GetStateStorage(config)
			state, err := stateStorage.Get()

			if err == nil && state.Project != "" {
				fmt.Println(chronolib.FormatStartError(state))
				return
			}
			project, tags, err := ParseStartArguments(args)
			if err != nil {
				PrintErrorAndExit(err)
			}
			newState, err := ParseNewFrameFlags(project, tags, startAt, startNote)
			if err != nil {
				PrintErrorAndExit(err)
			}
			newState, err = stateStorage.Update(newState)
			if err != nil {
				PrintErrorAndExit(err)
			}
			fmt.Println(chronolib.FormatNewFrameMessage(newState))
		},
	}
	startCmd.Flags().StringVarP(&startNote, "note", "n", "", "add an initial note to the frame")
	startCmd.Flags().StringVarP(&startAt, "at", "a", "", "set the start time to a different time than now - format: HH:MM mm/dd/yyyy")
	startCmd.Flags().StringVarP(&startEnded, "ended", "e", "", "add a manual end time to the new frame - does not get tracked through a timer")
	return startCmd
}
