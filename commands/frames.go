package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

var framesDescribe bool

func newFramesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "frames",
		Short: "Displays a list of all frame IDs",
		Long:  "Displays a list of all frame IDs",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			frames, _ := chronolib.GetFrames(config)
			for _, frame := range frames.All() {
				if framesDescribe {
					fmt.Println(chronolib.FormatFrameDescribe(frame))
				} else {
					fmt.Println(chronolib.GetShortHex(frame.UUID))
				}
			}
		},
	}
	cmd.Flags().BoolVarP(&framesDescribe, "describe", "d", false, "show information about frames instead of just their ID")
	return cmd
}
