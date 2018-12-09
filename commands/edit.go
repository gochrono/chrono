package commands

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gochrono/chrono/chronolib"

	"github.com/spf13/cobra"
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
            frameStorage := chronolib.GetFrameStorage(config)
            frames, err := frameStorage.All(chronolib.FrameFilterOptions{})
            if err != nil {
                commandError = err
                return
            }
			chronolib.SortFramesByDate(frames)

			var target chronolib.Frame
			var targetIndex int
			index, err := strconv.Atoi(args[0])
			if err != nil {
				targetIndex, target, err = chronolib.GetFrameByShortHex(frames, args[0])
				if err != nil {
					fmt.Println("No frame found with that ID")
					os.Exit(-1)
				}
			} else {
				if index < 0 {
					targetIndex = len(frames) + index
				} else {
					targetIndex = len(frames) - index
				}

                if targetIndex >= len(frames) || targetIndex < 0 {
					fmt.Println("No frame found at that index")
					os.Exit(-1)
                }
				target = frames[targetIndex]
			}

			simpleFrame := chronolib.ConvertFrameToSimpleFrame(target)
			simpleFrameJSON, err := json.MarshalIndent(simpleFrame, "", "    ")
			if err != nil {
				panic(err)
			}

			fpath := os.TempDir() + "/chrono-" + hex.EncodeToString(target.UUID) + ".json"
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
			frameEdited, err := chronolib.ConvertSimpleFrameToFrame(target.UUID, newSimpleFrame)
            if err != nil {
                fmt.Println(err)
                os.Exit(-1)
            }
			if chronolib.FramesEqual(target, frameEdited) {
				fmt.Println("No changes made")
			} else {
                frameStorage.Update(frameEdited)
				fmt.Println(chronolib.FormatEditFrameMessage(frameEdited))
			}
		},
	}
}
