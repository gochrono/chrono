package commands

import (
    "github.com/gochrono/chrono/chronolib"
)
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