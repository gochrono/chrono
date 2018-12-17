package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"time"
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
			jww.INFO.Printf("using configDir %s", configDir)
			config := chronolib.GetConfig(configDir)
			frames, _ := chronolib.GetFrames(config)
			state, _ := chronolib.GetState(config)

			jww.INFO.Printf("no argument, retrieving last frame")
			lastFrame, ok := frames.GetByIndex(-1)
			jww.DEBUG.Printf("last frame %v", lastFrame)

			if ok {
				notes := []string{}
				if restartNote != "" {
					notes = append(notes, restartNote)
				}
				currentFrame := chronolib.CurrentFrame{
					Project:   lastFrame.Project,
					StartedAt: time.Now(),
					UpdatedAt: time.Now(),
					Tags:      lastFrame.Tags,
					Notes:     notes,
				}
				state.Update(currentFrame)
				err := chronolib.SaveState(config, state)
				if err != nil {
					panic(err)
				}
				fmt.Println(chronolib.FormatStartCurrentFrame(currentFrame))
			} else {
				fmt.Println("Could not find frame")
			}
		},
	}
	cmd.Flags().StringVarP(&restartNote, "note", "n", "", "add an initial note to the frame")
	cmd.Flags().StringVarP(&restartAt, "at", "a", "", "set the start time to a different time than now - format: HH:MM mm/dd/yyyy")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return cmd
}
