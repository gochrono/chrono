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
			statePath := chronolib.GetAppFilePath("state", "")
			framesPath := chronolib.GetAppFilePath("frames", "")

			stateStorage := chronolib.GetStateStorage()
			// frameStorage := chronolib.GetFrameStorage()

			frame, err := stateStorage.Get()
			if err != nil {
				panic(err)
			}

			if stopAt != "" {
				t, err := now.Parse(stopAt)
				frame.EndedAt = t
				if err != nil {
					panic(err)
				}
			} else {
				frame.EndedAt = time.Now()
			}
			frame.UUID = chronolib.CreateFrameUUID(frame.Project, &frame.StartedAt, &frame.EndedAt)

			if stopNote != "" {
				frame.Notes = append(frame.Notes, stopNote)
			}

			data := chronolib.LoadFrames(framesPath)
			data.Frames = append(data.Frames, frame)
			chronolib.SaveFrames(framesPath, data)

			emptyFrame := chronolib.Frame{}
			chronolib.SaveState(statePath, &emptyFrame)

			fmt.Println(chronolib.FormatStopFrameMessage(frame))
		},
	}
	stopCmd.Flags().StringVarP(&stopNote, "note", "n", "", "add a final note to current frame")
	stopCmd.Flags().StringVarP(&stopAt, "at", "a", "", "sets the time the current frame ended to something other than now - format: HH:MM mm/dd/yyyy")
	return stopCmd
}
