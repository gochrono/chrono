package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

func newCancelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "Stops current project without saving it",
		Long:  "Stops current project without saving it",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			state, err := chronolib.GetState(config)
			if err != nil {
				panic(err)
			}

			if state.IsEmpty() {
				fmt.Println(chronolib.FormatNoProjectMessage())
			} else {
				fmt.Println(chronolib.FormatCancelMessage(state.Get()))
				state.Clear()
				err := chronolib.SaveState(config, state)
				if err != nil {
					panic(err)
				}
			}
		},
	}
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return cmd
}
