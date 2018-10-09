package commands

import (
    "os"
    "fmt"
    "time"
    "github.com/gochrono/chrono/chronolib"
    "github.com/spf13/cobra"
)

var forCurrentWeek bool
var forCurrentMonth bool
var forCurrentYear bool
var forAllTime bool



func newLogCmd() *cobra.Command {
    cmd :=  &cobra.Command{
        Use: "log",
        Short: "Get the frames logged for a single day",
        Long: "Get the frames logged for a single day",
        Run: func(cmd *cobra.Command, args []string) {
            framesPath := chronolib.GetAppFilePath("frames", "")
            data := chronolib.LoadFrames(framesPath)

            if chronolib.ContainsMoreThanOneBooleanFlag(forCurrentWeek, forCurrentMonth, forCurrentYear) {
                fmt.Println("Error: the folllowing flags are mutually exclusive: ['--week', '--year', '--month']")
                os.Exit(0)
            }

            var tsStart, tsEnd time.Time

            if forCurrentWeek {
                tsStart, tsEnd = chronolib.GetTimespanForWeek()
            } else if forCurrentMonth {
                tsStart, tsEnd = chronolib.GetTimespanForMonth()
            } else if forCurrentYear {
                tsStart, tsEnd = chronolib.GetTimespanForYear()
            } else {
                tsStart, tsEnd = chronolib.GetTimespanForToday()
            }

            filteredFrames := chronolib.FilterFramesByTimespan(tsStart, tsEnd, &data.Frames, forAllTime)
            dates := chronolib.SortTimeMapKeys(&filteredFrames)
            for _, date := range dates {
                fmt.Println(chronolib.FormatDateHeader(date))
                for _, frame := range filteredFrames[date] {
                    fmt.Println(chronolib.FormatFrameLine(frame))
                    for _, note := range frame.Notes {
                        fmt.Println(chronolib.FormatNoteLine(note))
                    }
                }
            }
        },
    }
    cmd.Flags().BoolVarP(&forCurrentWeek, "week", "w", false, "show frames for entire week")
    cmd.Flags().BoolVarP(&forCurrentMonth, "month", "m", false, "show frames for entire month")
    cmd.Flags().BoolVarP(&forCurrentYear, "year", "y", false, "show frames for entire year")
    cmd.Flags().BoolVarP(&forAllTime, "all", "a", false, "show all frames")
    return cmd
}
