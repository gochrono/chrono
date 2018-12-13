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
			state, _ := chronolib.GetState(config)

			if !state.IsEmpty() {
				fmt.Println(state.Get())
			}

			currentFrame, ended, err := ParseStartArgs(args, startAt, startEnded, startNote)
			if err != nil {
				PrintErrorAndExit(err)
			}

			state.Update(currentFrame)

			if startEnded != "" {
				frame := state.ToFrame(ended)
				frames, err := chronolib.GetFrames(config)
				if err != nil {
					panic(err)
				}
				frames.Add(frame)
				err = chronolib.SaveFrames(config, frames)
				if err != nil {
					panic(err)
				}
				state.Clear()
				fmt.Println(chronolib.FormatStartFrame(frame))
			} else {
				fmt.Println(chronolib.FormatStartCurrentFrame(state.Get()))
			}

			chronolib.SaveState(config, state)
		},
	}
	startCmd.Flags().StringVarP(&startNote, "note", "n", "", "add an initial note to the frame")
	startCmd.Flags().StringVarP(&startAt, "at", "a", "", "set the start time to a different time than now - format: HH:MM mm/dd/yyyy")
	startCmd.Flags().StringVarP(&startEnded, "ended", "e", "", "add a manual end time to the new frame - does not get tracked through a timer")
	startCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return startCmd
}
