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

func newReportCmd() *cobra.Command {
	newReport := &cobra.Command{
		Use:   "report",
		Short: "Get the total time spent on projects",
		Long:  "Get the total time spent on projects",
		Run: func(cmd *cobra.Command, args []string) {
			framesPath := chronolib.GetAppFilePath("frames", "")
			data := chronolib.LoadFrames(framesPath)

			if chronolib.ContainsMoreThanOneBooleanFlag(reportForCurrentWeek, reportForCurrentMonth, reportForCurrentYear, reportForAllTime) {
				fmt.Println("Error: the folllowing flags are mutually exclusive: ['--week', '--year', '--month', `--all`]")
				os.Exit(0)
			}

			var tsStart, tsEnd time.Time

			if reportForCurrentWeek {
				tsStart, tsEnd = chronolib.GetTimespanForWeek()
			} else if reportForCurrentMonth {
				tsStart, tsEnd = chronolib.GetTimespanForMonth()
			} else if reportForCurrentYear {
				tsStart, tsEnd = chronolib.GetTimespanForYear()
			} else {
				tsStart, tsEnd = chronolib.GetTimespanForToday()
			}

			filteredFrames := chronolib.FilterFramesByTimespan(tsStart, tsEnd, &data.Frames, reportForAllTime)
			dates := chronolib.SortTimeMapKeys(&filteredFrames)
			totals := make(map[string]frameTotals)
			fmt.Println(chronolib.FormatReportDuration(tsStart))
			for _, date := range dates {
				for _, frame := range filteredFrames[date] {
					frameTotal, ok := totals[frame.Project]
					frameDuration := frame.EndedAt.Sub(frame.StartedAt)
					if ok {
						frameTotal.TotalTime = frameTotal.TotalTime + frameDuration
					} else {
						totals[frame.Project] = frameTotals{frameDuration, make(map[string]time.Duration)}
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
	return newReport
}
