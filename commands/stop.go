package commands

import (
    "fmt"
    "time"
    "github.com/jordanknott/chrono/chronolib"
    "github.com/spf13/cobra"
)

func newStopCmd() *cobra.Command {
    return &cobra.Command{
        Use: "stop",
        Short: "Stop the current frame",
        Long: "Stop the current frame",
        Run: func(cmd *cobra.Command, args []string) {
            statePath := chronolib.GetAppFilePath("state", "")
            framesPath := chronolib.GetAppFilePath("frames", "")

            frame := chronolib.LoadState(statePath)
            now := time.Now()
            frame.EndedAt = now
            frame.UUID = chronolib.CreateFrameUUID(frame.Project, &frame.StartedAt, &frame.EndedAt)

            data := chronolib.LoadFrames(framesPath)
            data.Frames = append(data.Frames, *frame)
            chronolib.SaveFrames(framesPath, data)

            emptyFrame := chronolib.Frame{}
            chronolib.SaveState(statePath, &emptyFrame)

            fmt.Println(chronolib.FormatStopFrameMessage(*frame))
        },
    }
}
