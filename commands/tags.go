package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/jordanknott/chrono/chronolib"
)

func newTagsCmd() *cobra.Command {
    return &cobra.Command{
        Use: "tags",
        Short: "Get a list of all tags used",
        Long: "Get a list of all tags used",
        Run: func(cmd *cobra.Command, args []string) {
            framesPath := chronolib.GetAppFilePath("frames", "")
            data := chronolib.LoadFrames(framesPath)

            encounted := map[string]bool{}

            for _, frame := range data.Frames {
                for _, tag := range frame.Tags {
                    encounted[tag] = true
                }
            }

            for key, _ := range encounted {
                fmt.Println(key)
            }
        },
    }
}
