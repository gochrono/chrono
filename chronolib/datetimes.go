package chronolib

import (
	"github.com/SaidinWoT/timespan"
	"github.com/jinzhu/now"
	jww "github.com/spf13/jwalterweatherman"
	"time"
)

// ExtendTimeFormats adds support for various different time formats
func ExtendTimeFormats() {
	now.TimeFormats = append(now.TimeFormats, "02/01/2006 15:04") // dd/mm/yyyy HH:MM
}

// GetTimeElapsedForDuration returns the hours, minutes, seconds for a given duration
func GetTimeElapsedForDuration(delta time.Duration) (int, int, int) {
	hours := int(delta.Hours())
	minutes := int(delta.Minutes()) - (hours * 60)
	seconds := int(delta.Seconds()) - (hours * 60 * 60) - (minutes * 60)
	return hours, minutes, seconds
}

// GetTimeElapsed calculates the delta between two Times and returns the hours, minutes, and seconds for that delta
func GetTimeElapsed(dateStart, dateEnd time.Time) (int, int, int) {
	delta := dateEnd.Sub(dateStart)
	return GetTimeElapsedForDuration(delta)
}

// IsTimespanNegative returns true if a timespan has a negative amount of time, false otherwise
func IsTimespanNegative(start time.Time, end time.Time) bool {
	return end.Sub(start) < 0
}

// NormalizeDate strips hours, minutes, and seconds from a given time
func NormalizeDate(precise time.Time) time.Time {
	return time.Date(precise.Year(), precise.Month(), precise.Day(), 0, 0, 0, 0, precise.Location())
}

// SanitizeUserInput parses user-supplied date into specified format and
// returns time.Now() if date cannot be parsed.
func SanitizeUserInput(date string, format string) (time.Time) {
	parsed, err := time.Parse(format, date)
	if err != nil {
		jww.ERROR.Printf("couldn't parse specified date (%s), using default instead [err: %s]", date, err)
		return time.Now()
	}
	return parsed
}

// GetTimespanForToday returns the the start date and end date of today
func GetTimespanForToday() (time.Time, time.Time) {
	return now.BeginningOfDay(), now.EndOfDay()
}

// GetTimespanForWeek returns the start and end date for the current week
// TODO: Is a 2006-01-01 date format the best way to specify a week?
func GetTimespanForWeek(date string) (time.Time, time.Time) {
	parsed := SanitizeUserInput(date, "2006-01-02")
	return now.With(parsed).BeginningOfWeek(), now.With(parsed).EndOfWeek()
}

// GetTimespanForMonth returns the start and end date for the month passed in parameter
// with the format "2019-11"
func GetTimespanForMonth(month string) (time.Time, time.Time) {
	parsed := SanitizeUserInput(month, "2006-01")
	return now.With(parsed).BeginningOfMonth(), now.With(parsed).EndOfMonth()
}

// GetTimespanForDay returns the start and end date for specific day
func GetTimespanForDay(day string) (time.Time, time.Time) {
	parsed := SanitizeUserInput(day, "2006-01-02")
	return now.With(parsed).BeginningOfDay(), now.With(parsed).EndOfDay()
}

// GetTimespanForYear returns the start and end date for the current year
func GetTimespanForYear(year string) (time.Time, time.Time) {
	parsed := SanitizeUserInput(year, "2006")
	return now.With(parsed).BeginningOfYear(), now.With(parsed).EndOfYear()
}

// IsTimeInTimespan checks if point is inside the timespan between start and end
func IsTimeInTimespan(point time.Time, start time.Time, end time.Time) bool {
	duration := end.Sub(start)
	ts := timespan.New(start, duration)
	return ts.ContainsTime(point)
}

// HasSameDate checks if the date (month, year, day) are the same
func HasSameDate(t1 *time.Time, t2 *time.Time) bool {
	if t1.Day() == t2.Day() && t1.Year() == t2.Year() && t1.Month() == t2.Month() {
		return true
	}
	return false
}

// FilterFrames filters out a list of frames based on the given FrameFilterOptions
func FilterFrames(frames *[]Frame, filterOptions FrameFilterOptions) []Frame {
	var filteredFrames = []Frame{}
	var validFrame bool
	var useTags = len(filterOptions.Tags) != 0
	var useProjects = len(filterOptions.Projects) != 0
	var noTimespanCheck = filterOptions.TimespanFilter == TimespanFilterOptions{}
	var start = filterOptions.TimespanFilter.Start
	var end = filterOptions.TimespanFilter.End
	jww.DEBUG.Printf("FilterFrames start: %s", filterOptions.TimespanFilter.Start)
	jww.DEBUG.Printf("FilterFrames end: %s", filterOptions.TimespanFilter.End)
	for _, frame := range *frames {
		if IsFrameInTimespan(frame, start, end) || noTimespanCheck {
			validFrame = true
			if !IsWholeFrameInTimespan(frame, start, end) {
				frame = SubFrameForTimespan(frame, start, end)
			}
			if useTags {
				for _, tag := range filterOptions.Tags {
					if !StringInSlice(tag, frame.Tags) {
						validFrame = false
						break
					}
				}
			}
			if useProjects {
				if !StringInSlice(frame.Project, filterOptions.Projects) {
					validFrame = false
				}
			}
			if validFrame {
				filteredFrames = append(filteredFrames, frame)
			}
		}
	}

	return filteredFrames
}

// OrganizeFrameByTime returns a map of frames where the key is the date of the frame
func OrganizeFrameByTime(frames *[]Frame) map[time.Time][]Frame {
	frameMap := make(map[time.Time][]Frame)
	for _, frame := range *frames {
		date := NormalizeDate(frame.StartedAt)
		frameMap[date] = append(frameMap[date], frame)
	}
	return frameMap
}

// SubFrameForTimespan returns a new frame removing the part of the given frame that's outside the given timespan
func SubFrameForTimespan(frame Frame, start time.Time, end time.Time) Frame {
	newFrame := frame
	if !IsTimeInTimespan(frame.StartedAt, start, end) {
		newFrame.StartedAt = start
	}
	if !IsTimeInTimespan(frame.EndedAt, start, end) {
		newFrame.EndedAt = end
	}
	return newFrame
}

// IsWholeFrameInTimespan checks if a frame's start and end time are both in the given timespan
func IsWholeFrameInTimespan(frame Frame, start time.Time, end time.Time) bool {
	if !IsTimeInTimespan(frame.StartedAt, start, end) {
		return false
	}
	if !IsTimeInTimespan(frame.EndedAt, start, end) {
		return false
	}
	return true
}

// IsFrameInTimespan checks if a frame's start OR end time are in the given timespan
func IsFrameInTimespan(frame Frame, start time.Time, end time.Time) bool {
	return IsTimeInTimespan(frame.StartedAt, start, end) || IsTimeInTimespan(frame.EndedAt, start, end)
}
