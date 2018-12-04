package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var interval int
var forCurrentWeek bool
var forCurrentMonth bool
var forCurrentYear bool
var forAllTime bool
var round bool
var logTags []string

func newLogCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log",
		Short: "Get the frames logged for a single day",
		Long:  "Get the frames logged for a single day",
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

			filteredFrames := chronolib.FilterFramesByTimespan(tsStart, tsEnd, &data.Frames, forAllTime, logTags)
			dates := chronolib.SortTimeMapKeys(&filteredFrames)
			for _, date := range dates {
				fmt.Println(chronolib.FormatDateHeader(date))
				for _, frame := range filteredFrames[date] {
					if round {
						newStartTime := GetAdjustedTime(frame.StartedAt)
						newEndTime := GetAdjustedTime(frame.EndedAt)
						if !chronolib.IsTimespanNegative(newStartTime, newEndTime) {
							frame.StartedAt = newStartTime
							frame.EndedAt = newEndTime
						}
						fmt.Println(chronolib.FormatFrameLine(frame))
					} else {
						fmt.Println(chronolib.FormatFrameLine(frame))
					}
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
	cmd.Flags().BoolVarP(&round, "round", "r", false, "round frames start and end times to the nearest interval (default: 5 mins)")
	cmd.Flags().IntVarP(&interval, "interval", "i", 5, "the interval to round to in minutes")
	cmd.Flags().StringSliceVarP(&logTags, "tag", "t", []string{}, "only show frames that contain the given tag - can be used multiple times")
	return cmd
}

// GetAdjustedTime rounds a time to a given interval.
// E.g if the time is 15:02:13 and the interval is 5 minutes it
// would be rounded to 15:00:00
func GetAdjustedTime(t time.Time) time.Time {
	halfway := interval * 60 / 2
	rem := t.Minute()&interval*60 + t.Second()
	minutes := t.Minute() % interval
	if rem > halfway {
		if minutes == 0 {
			return t
		}
		return t.Add(time.Duration(interval-minutes) * time.Minute)
	}
	return t.Add(time.Duration(-minutes) * time.Minute)
}
