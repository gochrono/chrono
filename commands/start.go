package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
	"time"
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
			stateStorage := chronolib.GetStateStorage()
			state, err := stateStorage.Get()

			if err == nil && state.Project != "" {
				fmt.Println(chronolib.FormatStartError(state))
				return
			}

			project, tags, err := chronolib.ParseStartArguments(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			frameStart, err := chronolib.ParseTime(startAt)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}

			notes := []string{}
			if startNote != "" {
				notes = append(notes, startNote)
			}

			newState := chronolib.Frame{
				UUID: []byte{}, Project: project, StartedAt: frameStart, EndedAt: time.Time{}, Tags: tags, Notes: notes}

			newState, err = stateStorage.Update(newState)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}

			fmt.Println(chronolib.FormatNewFrameMessage(newState))
		},
	}
	startCmd.Flags().StringVarP(&startNote, "note", "n", "", "add an initial note to the frame")
	startCmd.Flags().StringVarP(&startAt, "at", "a", "", "set the start time to a different time than now - format: HH:MM mm/dd/yyyy")
	startCmd.Flags().StringVarP(&startEnded, "ended", "e", "", "add a manual end time to the new frame - does not get tracked through a timer")
	return startCmd
}
