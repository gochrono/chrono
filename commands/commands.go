package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

const mainDescription = `Chrono is a time to help track what you spend your time on.

You can start tracking your time with ` + "`start`" + ` and you can
stop the timer with ` + "`stop`" + `.`

var commandError error

// Execute creates the root command with all sub-commands installed
func Execute() {
	var rootCmd = &cobra.Command{
		Use:  "chrono",
		Long: mainDescription,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if commandError != nil {
				switch commandError.(type) {
				case *chronolib.ErrFileDoesNotExist:
					fmt.Println("error: " + commandError.Error())
				case *chronolib.ErrStateFileDoesNotExist:
					fmt.Println(chronolib.FormatNoProjectMessage())
				case *chronolib.ErrFramesFileDoesNotExist:
					fmt.Println("No frames exist")
				default:
					panic(commandError)
				}
			}
		},
	}
	rootCmd.AddCommand(newStartCmd(), newStatusCmd(), newStopCmd(), newReportCmd(),
		newLogCmd(), newEditCmd(), newVersionCmd(), newNotesCmd(), newTagsCmd())
	rootCmd.Execute()
}
