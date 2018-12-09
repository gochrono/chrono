package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

func newCancelCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cancel",
		Short: "Stops current project without saving it",
		Long:  "Stops current project without saving it",
		Run: func(cmd *cobra.Command, args []string) {
            configDir := chronolib.GetCorrectConfigDirectory("")
            config := chronolib.GetConfig(configDir)
			stateStorage := chronolib.GetStateStorage(config)
            state, err := stateStorage.Get()
            if err != nil {
                commandError = err
                return
            }
            _, err = stateStorage.Clear()
            if err != nil {
                commandError = err
                return
            }
            fmt.Println(chronolib.FormatCancelMessage(state))
		},
	}
}
