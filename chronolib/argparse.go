package chronolib

import (
	"errors"
	"github.com/jinzhu/now"
	"time"
)

// ParseStartArguments splits the argument string list and validates tags
func ParseStartArguments(args []string) (string, []string, error) {
	project := args[0]
	tags := args[1:]

	if err := CheckTags(tags); err != nil {
		return "", []string{}, err
	}

	return project, NormalizeTags(tags), nil
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
