package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

var restartAt string
var restartNote string

func newRestartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart",
		Short: "Restart time tracking for a previously stopped project",
		Long:  "Restart time tracking for a previously stopped project",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			frameStorage := chronolib.GetFrameStorage(config)
			stateStorage := chronolib.GetStateStorage(config)
			lastFrame, err := frameStorage.Get(chronolib.FrameGetOptions{Target: "-1"})
			if err != nil {
				PrintErrorAndExit(err)
			}
			newState, err := ParseNewFrameFlags(
				lastFrame.Project,
				lastFrame.Tags,
				restartAt,
				restartNote,
			)
			if err != nil {
				PrintErrorAndExit(err)
			}
			_, err = stateStorage.Update(newState)
			if err != nil {
				PrintErrorAndExit(err)
			}
			fmt.Println(chronolib.FormatNewFrameMessage(newState))
		},
	}
	cmd.Flags().StringVarP(&restartNote, "note", "n", "", "add an initial note to the frame")
	cmd.Flags().StringVarP(&restartAt, "at", "a", "", "set the start time to a different time than now - format: HH:MM mm/dd/yyyy")
	return cmd
}
