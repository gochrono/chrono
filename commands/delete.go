package commands

import (
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
    "fmt"
)


func newDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Deletes a frame through either its index possition or UUID",
		Long:  "Deletes a frame through either its index possition or UUID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
            configDir := chronolib.GetCorrectConfigDirectory("")
            config := chronolib.GetConfig(configDir)
            frameStorage := chronolib.GetFrameStorage(config)
            frame, err := frameStorage.Delete(chronolib.FrameDeleteOptions{
                Target: args[0],
            })
            if err != nil {
                panic(err)
            }
            fmt.Println(frame)
		},
	}
}
