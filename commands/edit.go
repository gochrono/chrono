package commands

import (
	"encoding/json"
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

const editDesciption = `Edit a frame.

You can specific which frame to edit through either its ID or relative position in the index.
For example if you wanted to get the last frame you edited
you would do ` + "`chrono edit -- -1`" + `.

The '--' is needed so that the negative number is not parsed
as a flag.

If you wanted to get the earliest frame,
you would do ` + "`chrono edit 1`" + `.`

func newEditCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "edit [index]",
		Short: "Edit a frame",
		Long:  editDesciption,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			frames, _ := chronolib.GetFrames(config)
			frame, ok := GetFrame(frames, args[0])
			if !ok {
				fmt.Println("Could not find frame")
				return
			}

			simpleFrame := chronolib.ConvertFrameToSimpleFrame(frame)
			simpleFrameJSON, err := json.MarshalIndent(simpleFrame, "", "    ")
			if err != nil {
				panic(err)
			}

			fpath := os.TempDir() + "/chrono-" + frame.UUID + ".json"
			err = ioutil.WriteFile(fpath, simpleFrameJSON, 0644)
			if err != nil {
				panic(err)
			}

			editor := chronolib.GetEditorEnv()
			chronolib.EditFile(editor, fpath)

			content, err := ioutil.ReadFile(fpath)
			if err != nil {
				panic(err)
			}

			var newSimpleFrame chronolib.SimpleFrame
			err = json.Unmarshal(content, &newSimpleFrame)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			frameEdited, err := chronolib.ConvertSimpleFrameToFrame(frame.UUID, newSimpleFrame)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			if chronolib.FramesEqual(frame, frameEdited) {
				fmt.Println("No changes made")
			} else {
				frames.Update(frameEdited)
				if err := chronolib.SaveFrames(config, frames); err != nil {
					panic(err)
				}
				fmt.Println(chronolib.FormatEditFrameMessage(frameEdited))
			}
		},
	}
}
