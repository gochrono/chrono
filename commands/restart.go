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
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			jww.INFO.Printf("using configDir %s", configDir)
			config := chronolib.GetConfig(configDir)
			frames, err := chronolib.GetFrames(config)
			if err != nil {
				panic(err)
			}
			state, err := chronolib.GetState(config)
			if err != nil {
				panic(err)
			}

			if !state.IsEmpty() {
				fmt.Println(chronolib.FormatStartError(state.Get()))
				return
			}

			var target string
			if len(args) == 0 {
				target = "-1"
				jww.INFO.Printf("no argument, retrieving last frame")
			} else {
				target = args[0]
				jww.INFO.Printf("using target %s", target)
			}

			lastFrame, ok := GetFrame(frames, target)

			if ok {
				jww.DEBUG.Printf("found last frame %v", lastFrame)
				notes := []string{}
				startedAt := time.Now()
				if restartNote != "" {
					notes = append(notes, restartNote)
				}
				if restartAt != "" {
					startedAt, err = ParseTime(restartAt)
					if err != nil {
						fmt.Println(chronolib.FormatTimeStringNotValid())
						return
					}
				}
				currentFrame := chronolib.CurrentFrame{
					Project:   lastFrame.Project,
					StartedAt: startedAt,
					UpdatedAt: startedAt,
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
	cmd.Flags().StringVarP(&restartAt, "at", "a", "", "set the start time to a different time than now - format: 'mm/dd/yyyy HH:MM'. Make sure it is wrapped in quotes.")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return cmd
}
