package chronolib

import (
    "time"
    "github.com/SaidinWoT/timespan"
    "github.com/jinzhu/now"
)

var FirstDayOfWeek = "Monday"


func GetTimeElapsedForDuration(delta time.Duration) (int, int, int) {
    hours := int(delta.Hours())
    minutes := int(delta.Minutes()) - (hours * 60)
    seconds := int(delta.Seconds()) - (hours * 60 * 60) - (minutes * 60)
    return hours, minutes, seconds
}

func GetTimeElapsed(dateStart, dateEnd time.Time) (int, int, int) {
    delta := dateEnd.Sub(dateStart)
    return GetTimeElapsedForDuration(delta)
}

func NormalizeDate(precise time.Time) time.Time {
    return time.Date(precise.Year(), precise.Month(), precise.Day(), 0, 0, 0, 0, precise.Location())
}

func GetTimespanForToday() (time.Time, time.Time) {
    return now.BeginningOfDay(), now.EndOfDay()
}

func GetTimespanForWeek() (time.Time, time.Time) {
    return now.BeginningOfWeek(), now.EndOfWeek()
}

func GetTimespanForMonth() (time.Time, time.Time) {
    return now.BeginningOfMonth(), now.EndOfMonth()
}

func GetTimespanForYear() (time.Time, time.Time) {
    return now.BeginningOfYear(), now.EndOfYear()
}

func IsTimeInTimespan(point time.Time, start time.Time, end time.Time) bool {
    duration := end.Sub(start)
    ts := timespan.New(start, duration)
    return ts.ContainsTime(point)
}

func HasSameDate(t1 *time.Time, t2 *time.Time) bool {
    if t1.Day() == t2.Day() && t1.Year() == t2.Year() && t1.Month() == t2.Month() {
        return true
    }
    return false
}

func FilterFramesByTimespan(start time.Time, end time.Time, frames *[]Frame, noCheck bool) map[time.Time][]Frame {
    filteredFrames := make(map[time.Time][]Frame)
    for _, frame := range *frames {
        if IsFrameInTimespan(frame, start, end) || noCheck {
            date := NormalizeDate(frame.StartedAt)
            filteredFrames[date] = append(filteredFrames[date], frame)
        }
    }
    return filteredFrames
}

func IsFrameInTimespan(frame Frame, start time.Time, end time.Time) bool {
    if !IsTimeInTimespan(frame.StartedAt, start, end) {
        return false
    }
    if !IsTimeInTimespan(frame.EndedAt, start, end) {
        return false
    }
    return true
}
