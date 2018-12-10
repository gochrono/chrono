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
			frameStorage := chronolib.GetFrameStorage(config)
			frames, err := frameStorage.All(chronolib.FrameFilterOptions{})
			if err != nil {
				commandError = err
				return
			}
			for _, frame := range frames {
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
