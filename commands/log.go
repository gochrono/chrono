package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/jinzhu/now"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"os"
	"time"
)

var interval int
var logForCurrentWeek bool
var logForCurrentMonth bool
var logForCurrentYear bool
var logForAllTime bool
var logForYesterday bool
var logFrom string
var logTo string
var round bool
var logTags []string
var logProjects []string

// GetToFromTimespan is a helper function that takes in two time formats and attempts to parse them
func GetToFromTimespan(from string, to string) (chronolib.TimespanFilterOptions, error) {
	start := now.BeginningOfDay()
	end := now.EndOfDay()
	var err error
	if logTo != "" {
		start, err = ParseTime(logTo)
		if err != nil {
			fmt.Println("unable to parse time string")
			return chronolib.TimespanFilterOptions{}, err
		}
	}
	if logFrom != "" {
		end, err = ParseTime(logFrom)
		if err != nil {
			fmt.Println("unable to parse time string")
			return chronolib.TimespanFilterOptions{}, err
		}
	}
	return chronolib.TimespanFilterOptions{Start: start, End: end}, nil
}

func newLogCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log",
		Short: "Get the frames logged for a single day",
		Long:  "Get the frames logged for a single day",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)

			if chronolib.ContainsMoreThanOneBooleanFlag(logForCurrentWeek, logForCurrentMonth, logForCurrentYear) {
				fmt.Println("Error: the folllowing flags are mutually exclusive: ['--week', '--year', '--month']")
				os.Exit(0)
			}

			var timespanFilterOptions chronolib.TimespanFilterOptions
			var err error
			if logTo != "" || logFrom != "" {
				timespanFilterOptions, err = GetToFromTimespan(logFrom, logTo)
				if err != nil {
					fmt.Println("unable to parse time string")
					return
				}
			} else {
				timespanFilterOptions = ParseTimespanFlags(TimespanFlags{
					AllTime:      logForAllTime,
					CurrentWeek:  logForCurrentWeek,
					CurrentMonth: logForCurrentMonth,
					CurrentYear:  logForCurrentYear,
					Yesterday:    logForYesterday,
				})
			}

			jww.INFO.Printf("timespan filter options: %v", timespanFilterOptions)

			frames, _ := chronolib.GetFrames(config)
			filteredFrames := frames.Filter(chronolib.FrameFilterOptions{
				TimespanFilter: timespanFilterOptions, Tags: logTags, Projects: logProjects,
			})
			timemap := chronolib.OrganizeFrameByTime(&filteredFrames)
			dates := chronolib.SortTimeMapKeys(&timemap)
			for _, date := range dates {
				fmt.Println(chronolib.FormatDateHeader(date))
				for _, frame := range timemap[date] {
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
	cmd.Flags().BoolVarP(&logForYesterday, "yesterday", "d", false, "show frames for yesterday")
	cmd.Flags().BoolVarP(&logForAllTime, "all", "a", false, "show all frames")
	cmd.Flags().BoolVarP(&round, "round", "r", false, "round frames start and end times to the nearest interval (default: 5 mins)")
	cmd.Flags().StringVarP(&logFrom, "from", "f", "", "")
	cmd.Flags().StringVarP(&logTo, "to", "T", "", "")
	cmd.Flags().IntVarP(&interval, "interval", "i", 5, "the interval to round to in minutes")
	cmd.Flags().StringSliceVarP(&logTags, "tag", "t", []string{}, "only show frames that contain the given tag - can be used multiple times")
	cmd.Flags().StringSliceVarP(&logProjects, "project", "p", []string{}, "only show frames that contain the given project - can be used multiple times")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
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
