package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

func newFramesCmd() *cobra.Command {
	return &cobra.Command{
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
                fmt.Println(chronolib.GetShortHex(frame.UUID))
			}
		},
	}
}
