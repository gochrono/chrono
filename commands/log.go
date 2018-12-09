package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var interval int
var logForCurrentWeek bool
var logForCurrentMonth bool
var logForCurrentYear bool
var logForAllTime bool
var round bool
var logTags []string

func newLogCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log",
		Short: "Get the frames logged for a single day",
		Long:  "Get the frames logged for a single day",
		Run: func(cmd *cobra.Command, args []string) {
            configDir := chronolib.GetCorrectConfigDirectory("")
            config := chronolib.GetConfig(configDir)
			frameStorage := chronolib.GetFrameStorage(config)

			if chronolib.ContainsMoreThanOneBooleanFlag(logForCurrentWeek, logForCurrentMonth, logForCurrentYear) {
				fmt.Println("Error: the folllowing flags are mutually exclusive: ['--week', '--year', '--month']")
				os.Exit(0)
			}

            timespanFilterOptions := ParseTimespanFlags(TimespanFlags{
                AllTime: logForAllTime,
                CurrentWeek: logForCurrentWeek,
                CurrentMonth: logForCurrentMonth,
                CurrentYear: logForCurrentYear,
            })

			frames, err := frameStorage.All(chronolib.FrameFilterOptions{
                TimespanFilter: timespanFilterOptions, Tags: logTags,
            })
			if err != nil {
				commandError = err
				return
			}

			filteredFrames := chronolib.OrganizeFrameByTime(&frames)
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
	cmd.Flags().BoolVarP(&logForCurrentWeek, "week", "w", false, "show frames for entire week")
	cmd.Flags().BoolVarP(&logForCurrentMonth, "month", "m", false, "show frames for entire month")
	cmd.Flags().BoolVarP(&logForCurrentYear, "year", "y", false, "show frames for entire year")
	cmd.Flags().BoolVarP(&logForAllTime, "all", "a", false, "show all frames")
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
