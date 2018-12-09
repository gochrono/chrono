package commands

import (
    "time"
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


func ParseNewFrameFlags(project string, tags []string, startAt string, startNote string) (chronolib.Frame, error) {
    frameStart, err := chronolib.ParseTime(startAt)
    if err != nil {
        return chronolib.Frame{}, err
    }

    notes := []string{}
    if startNote != "" {
        notes = append(notes, startNote)
    }

    return chronolib.Frame{
        UUID: []byte{}, Project: project, 
        StartedAt: frameStart, EndedAt: time.Time{}, 
        Tags: tags, Notes: notes,
        UpdatedAt: time.Now(),
    }, nil
}