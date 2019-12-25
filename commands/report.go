package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"time"
)

var reportForDay string
var reportForWeek string
var reportForMonth string
var reportForYear string
var reportForAllTime bool
var reportTags []string
var reportProjects []string

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
				reportForWeek != "", reportForMonth != "",
				reportForYear != "", reportForDay != "",
				reportForAllTime,
			) {
				fmt.Println("Error: the following flags are mutually exclusive: ['--day --week', '--year', '--month', `--all`]")
				os.Exit(0)
			}

			timespanFilterOptions := ParseTimespanFlags(TimespanFlags{
				AllTime:      reportForAllTime,
				Day:          reportForDay,
				Week:         reportForWeek,
				Month:        reportForMonth,
				Year:         reportForYear,
			})

			frames, _ := chronolib.GetFrames(config)
			filteredFrames := frames.Filter(chronolib.FrameFilterOptions{
				TimespanFilter: timespanFilterOptions,
				Tags:           reportTags,
				Projects:       reportProjects,
			})

			timemap := chronolib.OrganizeFrameByTime(&filteredFrames)
			dates := chronolib.SortTimeMapKeys(&timemap)
			totals := make(map[string]frameTotals)
			fmt.Println(chronolib.FormatReportDuration(
				timespanFilterOptions.Start,
				timespanFilterOptions.End,
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
				keys := make([]string, 0, len(frameTotal.Tags))
				for key := range frameTotal.Tags {
					keys = append(keys, key)
				}
				sort.Strings(keys)
				for _, tag := range keys {
					fmt.Println(chronolib.FormatReportProjectTagTotal(tag, frameTotal.Tags[tag]))
				}
			}
		},
	}

	newReport.Flags().StringVarP(&reportForWeek, "week", "w", "", "show frames for entire week")
	newReport.Flags().Lookup("week").NoOptDefVal = time.Now().Format("2006-01-02")
	newReport.Flags().StringVarP(&reportForMonth, "month", "m", "", "show frames for entire month")
	newReport.Flags().Lookup("month").NoOptDefVal = time.Now().Format("2006-01")
	newReport.Flags().StringVarP(&reportForYear, "year", "y", "", "show frames for entire year")
	newReport.Flags().Lookup("year").NoOptDefVal = time.Now().Format("2006")
	newReport.Flags().StringVarP(&reportForDay, "day", "d", "", "show frames for day")
	newReport.Flags().Lookup("day").NoOptDefVal = time.Now().Format("2006-01-02")
	newReport.Flags().BoolVarP(&reportForAllTime, "all", "a", false, "show all frames")
	newReport.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	newReport.Flags().StringSliceVarP(&reportTags, "tag", "t", []string{}, "only show frames that contain the given tag - can be used multiple times")
	newReport.Flags().StringSliceVarP(&reportProjects, "project", "p", []string{}, "only show frames that contain the given project - can be used multiple times")
	return newReport
}
