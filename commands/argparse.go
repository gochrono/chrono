package commands

import (
	"errors"
	"github.com/gochrono/chrono/chronolib"
	"github.com/jinzhu/now"
	"time"
)

// ParseStartArguments splits the argument string list and validates tags
func ParseStartArguments(args []string) (string, []string, error) {
	project := args[0]
	tags := args[1:]

	if err := chronolib.CheckTags(tags); err != nil {
		return "", []string{}, err
	}

	return project, chronolib.NormalizeTags(tags), nil
}

// ParseTime converts a properly formated time string into a time.Time struct
func ParseTime(t string) (time.Time, error) {
	if t == "" {
		return time.Now(), nil
	}
	parsedTime, err := now.Parse(t)
	if err != nil {
		return time.Time{}, errors.New("invalid time format: " + t)
	}
	return parsedTime, nil
}

// TimespanFlags is a struct containing the four different options for timespans
type TimespanFlags struct {
	AllTime      bool
	CurrentWeek  bool
	CurrentMonth bool
	CurrentYear  bool
}

// ParseTimespanFlags gets the correct start and end time for filtering frames based
// on the flags given
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

// ParseNewFrameFlags is a helper method for creating a new frame based on user input
func ParseNewFrameFlags(project string, tags []string, startAt string, startNote string) (chronolib.Frame, error) {
	frameStart, err := ParseTime(startAt)
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
