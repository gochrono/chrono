package commands

import (
    "os"
    "fmt"
    "strconv"
    "encoding/json"
    "encoding/hex"
    "io/ioutil"

    "github.com/jordanknott/chrono/chronolib"

    "github.com/spf13/cobra"
)

func newEditCmd() *cobra.Command {
    return &cobra.Command{
        Use: "edit [index]",
        Short: "edit a frame",
        Long: "edit a frame",
        Args: cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            framesPath := chronolib.GetAppFilePath("frames", "")
            data := chronolib.LoadFrames(framesPath)
            chronolib.SortFramesByDate(data.Frames)

            var target chronolib.Frame
            var targetIndex int
            index, err := strconv.Atoi(args[0])
            if err != nil {
                targetIndex, target, err = data.GetFrameByShortHex(args[0])
                if err != nil {
                    fmt.Println("No frame found with that ID")
                    os.Exit(-1)
                }
            } else {
                targetIndex = len(data.Frames) - index
                target, err = data.GetFrameByIndex(targetIndex)
                if err != nil {
                    fmt.Println("No frame found at that index")
                    os.Exit(-1)
                }
            }

            frameRaw := chronolib.ConvertFrameToRawFrame(target)
            frameRawJson, err := json.MarshalIndent(frameRaw, "", "    ")
            if err != nil { panic(err) }

            fpath := os.TempDir() + "/chrono-" + hex.EncodeToString(target.UUID) + ".json"
            err = ioutil.WriteFile(fpath, frameRawJson, 0644)

            editor := chronolib.GetEditorEnv()
            chronolib.EditFile(editor, fpath)

            content, err := ioutil.ReadFile(fpath)

            var frameRawEdited chronolib.FrameRaw
            err = json.Unmarshal(content, &frameRawEdited)
            if err != nil { panic(err) }
            frameEdited := chronolib.ConvertRawFrameToFrame(target.UUID, frameRawEdited)
            if chronolib.FramesEqual(target, frameEdited) {
                fmt.Println("No changes made")
            } else {
                data.Frames[targetIndex] = frameEdited
                chronolib.SaveFrames(framesPath, data)

                started := frameEdited.StartedAt.Format("15:04:05")
                ended := frameEdited.EndedAt.Format("15:04:05")
                hours, minutes, seconds := chronolib.GetTimeElapsed(frameEdited.StartedAt, frameEdited.EndedAt)
                fmt.Printf("Edited frame for project %s, from %s to %s (%dh %02dm %02ds)\n", frameEdited.Project, started, ended, hours, minutes, seconds)
            }
        },
    }
}
