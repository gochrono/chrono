package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var reportForCurrentWeek bool
var reportForCurrentMonth bool
var reportForCurrentYear bool
var reportForAllTime bool

type frameTotals struct {
	TotalTime time.Duration
	Tags      map[string]time.Duration
}

func (f *frameTotals) SetTotalTime(d time.Duration) {
	(*f).TotalTime = f.TotalTime + d
}

func newReportCmd() *cobra.Command {
	newReport := &cobra.Command{
		Use:   "report",
		Short: "Get the total time spent on projects",
		Long:  "Get the total time spent on projects",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)

			if chronolib.ContainsMoreThanOneBooleanFlag(
				reportForCurrentWeek, reportForCurrentMonth,
				reportForCurrentYear, reportForAllTime,
			) {
				fmt.Println("Error: the folllowing flags are mutually exclusive: ['--week', '--year', '--month', `--all`]")
				os.Exit(0)
			}

			timespanFilterOptions := ParseTimespanFlags(TimespanFlags{
				AllTime:      reportForAllTime,
				CurrentWeek:  reportForCurrentWeek,
				CurrentMonth: reportForCurrentMonth,
				CurrentYear:  reportForCurrentYear,
			})

			frames, _ := chronolib.GetFrames(config)
			filteredFrames := frames.Filter(chronolib.FrameFilterOptions{
				TimespanFilter: timespanFilterOptions, Tags: logTags,
			})

			timemap := chronolib.OrganizeFrameByTime(&filteredFrames)
			dates := chronolib.SortTimeMapKeys(&timemap)
			totals := make(map[string]frameTotals)
			fmt.Println(chronolib.FormatReportDuration(
				timespanFilterOptions.Start,
			))
			for _, date := range dates {
				for _, frame := range timemap[date] {
					_, ok := totals[frame.Project]
					frameDuration := frame.EndedAt.Sub(frame.StartedAt)
					if ok {
						totals[frame.Project] = frameTotals{
							totals[frame.Project].TotalTime + frameDuration, totals[frame.Project].Tags,
						}
					} else {
						totals[frame.Project] = frameTotals{
							frameDuration, make(map[string]time.Duration),
						}
					}

					for _, tag := range frame.Tags {
						totals[frame.Project].Tags[tag] = totals[frame.Project].Tags[tag] + frameDuration
					}
				}
			}

			for project, frameTotal := range totals {
				fmt.Println(chronolib.FormatReportProjectTotal(project, frameTotal.TotalTime))
				for tag, duration := range frameTotal.Tags {
					fmt.Println(chronolib.FormatReportProjectTagTotal(tag, duration))
				}
			}
		},
	}

	newReport.Flags().BoolVarP(&reportForCurrentWeek, "week", "w", false, "show frames for entire week")
	newReport.Flags().BoolVarP(&reportForCurrentMonth, "month", "m", false, "show frames for entire month")
	newReport.Flags().BoolVarP(&reportForCurrentYear, "year", "y", false, "show frames for entire year")
	newReport.Flags().BoolVarP(&reportForAllTime, "all", "a", false, "show all frames")
	newReport.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	return newReport
}
