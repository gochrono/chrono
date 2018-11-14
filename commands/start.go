package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/jinzhu/now"
	"github.com/spf13/cobra"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
	"time"
)

var startNote string
var startTime string
var startEndTime string

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
			statePath := chronolib.GetAppFilePath("state", "")

			var project = args[0]
			var tags = args[1:]
			var frameStart = time.Now()

			if !chronolib.IsAllTags(tags) {
				fmt.Println("Invalid tag")
				os.Exit(-1)
			}

			tags = chronolib.NormalizeTags(tags)

			var notes []string
			if startNote == "" {
				notes = []string{}
			} else {
				notes = []string{startNote}
			}

			if startTime != "" {
				t, err := now.Parse(startTime)
				frameStart = t
				if err != nil {
					panic(err)
				}
			}
			frame := chronolib.Frame{
				UUID: []byte{}, Project: project, StartedAt: frameStart, EndedAt: time.Time{}, Tags: tags, Notes: notes}

			b, err := msgpack.Marshal(&frame)
			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(statePath, b, 0644)
			if err != nil {
				panic(err)
			}

			fmt.Println(chronolib.FormatNewFrameMessage(frame))
		},
	}
	startCmd.Flags().StringVarP(&startNote, "note", "n", "", "add an initial note to the frame")
	startCmd.Flags().StringVarP(&startTime, "start", "s", "", "set the start time to a different time than now - format: HH:MM mm/dd/yyyy")
	startCmd.Flags().StringVarP(&startEndTime, "end", "e", "", "add a manual end time to the new frame - does not get tracked through a timer")
	return startCmd
}
