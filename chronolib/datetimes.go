package chronolib

import (
    "time"
    "github.com/SaidinWoT/timespan"
    "github.com/jinzhu/now"
)

var FirstDayOfWeek = "Monday"



func GetTimeElapsed(dateStart, dateEnd time.Time) (int, int, int) {
    delta := dateEnd.Sub(dateStart)
    hours := int(delta.Hours())
    minutes := int(delta.Minutes()) - (hours * 60)
    seconds := int(delta.Seconds()) - (hours * 60 * 60) - (minutes * 60)
    return hours, minutes, seconds
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

