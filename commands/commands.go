package commands

import (
	"fmt"
    "time"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

const mainDescription = `Chrono is a time to help track what you spend your time on.

You can start tracking your time with ` + "`start`" + ` and you can
stop the timer with ` + "`stop`" + `.`

type TimespanFlags struct {
    AllTime        bool
    CurrentWeek    bool
    CurrentMonth    bool
    CurrentYear    bool
}

func ParseTimespanFlags(timespanFlags TimespanFlags) chronolib.TimespanFilterOptions {
	var tsStart, tsEnd time.Time
    if timespanFlags.AllTime {
	    return chronolib.TimespanFilterOptions{}
    } else if timespanFlags.CurrentWeek {
	    tsStart, tsEnd = chronolib.GetTimespanForWeek()
    } else if timespanFlags.CurrentMonth {
		tsStart, tsEnd = chronolib.GetTimespanForMonth()
    } else if timespanFlags.CurrentYear {
		tsStart, tsEnd = chronolib.GetTimespanForYear()
    } else {
		tsStart, tsEnd = chronolib.GetTimespanForToday()
    }
	return chronolib.TimespanFilterOptions{Start: tsStart, End: tsEnd}
}

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
					fmt.Println(chronolib.FormatNoFramesMessage())
				default:
					panic(commandError)
				}
			}
		},
	}
	rootCmd.AddCommand(newStartCmd(), newStatusCmd(), newStopCmd(), newReportCmd(),
                       newLogCmd(), newCancelCmd(), newDeleteCmd(), newFramesCmd(),
                       newEditCmd(), newVersionCmd(), newNotesCmd(), newTagsCmd())
	rootCmd.Execute()
}
