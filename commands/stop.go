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
			stateStorage := chronolib.GetStateStorage(config)
			frameStorage := chronolib.GetFrameStorage(config)

			frame, err := stateStorage.Get()
			if err != nil {
				commandError = err
				return
			}

			if stopAt != "" {
				t, err := now.Parse(stopAt)
				frame.EndedAt = t
				if err != nil {
					commandError = err
					return
				}
			} else {
				frame.EndedAt = time.Now()
			}
			frame.UUID = chronolib.CreateFrameUUID(frame.Project, &frame.StartedAt, &frame.EndedAt)

			if stopNote != "" {
				frame.Notes = append(frame.Notes, stopNote)
			}

			_, err = frameStorage.Add(frame)
			_, err = stateStorage.Clear()

			fmt.Println(chronolib.FormatStopFrameMessage(frame))
		},
	}
	stopCmd.Flags().StringVarP(&stopNote, "note", "n", "", "add a final note to current frame")
	stopCmd.Flags().StringVarP(&stopAt, "at", "a", "", "sets the time the current frame ended to something other than now - format: HH:MM mm/dd/yyyy")
	return stopCmd
}
