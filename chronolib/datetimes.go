package chronolib

import (
	"github.com/SaidinWoT/timespan"
	"github.com/jinzhu/now"
	"time"
)

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

// GetTimespanForToday returns the the start date and end date of today
func GetTimespanForToday() (time.Time, time.Time) {
	return now.BeginningOfDay(), now.EndOfDay()
}

// GetTimespanForWeek returns the start and end date for the current week
func GetTimespanForWeek() (time.Time, time.Time) {
	return now.BeginningOfWeek(), now.EndOfWeek()
}

// GetTimespanForMonth returns the start and end date for the current month
func GetTimespanForMonth() (time.Time, time.Time) {
	return now.BeginningOfMonth(), now.EndOfMonth()
}

// GetTimespanForYear returns the start and end date for the current year
func GetTimespanForYear() (time.Time, time.Time) {
	return now.BeginningOfYear(), now.EndOfYear()
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
	var noTimespanCheck = filterOptions.TimespanFilter == TimespanFilterOptions{}
	var start = filterOptions.TimespanFilter.Start
	var end = filterOptions.TimespanFilter.End
	for _, frame := range *frames {
		if IsFrameInTimespan(frame, start, end) || noTimespanCheck {
			validFrame = true
			if useTags {
				for _, tag := range filterOptions.Tags {
					if !StringInSlice(tag, frame.Tags) {
						validFrame = false
						break
					}
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

// FilterFramesByTimespan returns only frames that are in the given timespan
func FilterFramesByTimespan(start time.Time, end time.Time, frames *[]Frame, noCheck bool, tags []string) map[time.Time][]Frame {
	filteredFrames := make(map[time.Time][]Frame)
	var validFrame bool
	var useTags = len(tags) != 0
	for _, frame := range *frames {
		if IsFrameInTimespan(frame, start, end) || noCheck {
			validFrame = true
			if useTags {
				for _, tag := range tags {
					if !StringInSlice(tag, frame.Tags) {
						validFrame = false
						break
					}
				}
			}
			if validFrame {
				date := NormalizeDate(frame.StartedAt)
				filteredFrames[date] = append(filteredFrames[date], frame)
			}
		}
	}
	return filteredFrames
}

// IsFrameInTimespan checks if a frame's start and end time are both in the given timespan
func IsFrameInTimespan(frame Frame, start time.Time, end time.Time) bool {
	if !IsTimeInTimespan(frame.StartedAt, start, end) {
		return false
	}
	if !IsTimeInTimespan(frame.EndedAt, start, end) {
		return false
	}
	return true
}
