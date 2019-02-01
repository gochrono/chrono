package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/jinzhu/now"
	"github.com/spf13/cobra"
	"time"
)

var stopNote string
var stopAt string

func newStopCmd() *cobra.Command {
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop the current frame",
		Long:  "Stop the current frame",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			state, _ := chronolib.GetState(config)

			if state.IsEmpty() {
				fmt.Println("No project started")
				return
			}

			var endedAt time.Time
			var err error
			if stopAt != "" {
				endedAt, err = now.Parse(stopAt)
				if err != nil {
					fmt.Println("Unable to parse time string")
					return
				}
			} else {
				endedAt = time.Now()
			}

			newFrame := state.ToFrame(endedAt)

			if stopNote != "" {
				newFrame.Notes = append(newFrame.Notes, stopNote)
			}

			frames, _ := chronolib.GetFrames(config)
			frames.Add(newFrame)
			chronolib.SaveFrames(config, frames)
			state.Clear()
			chronolib.SaveState(config, state)

			fmt.Println(chronolib.FormatStopFrameMessage(newFrame))
		},
	}
	stopCmd.Flags().StringVarP(&stopNote, "note", "n", "", "add a final note to current frame")
	stopCmd.Flags().StringVarP(&stopAt, "at", "a", "", "sets the time the current frame ended to something other than now - format: 'mm/dd/yyyy HH:MM'. Make sure it is wrapped in quotes.")
	stopCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return stopCmd
}
