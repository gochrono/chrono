package chronolib

import (
	"errors"
	"time"
)

var keywords = []string{
	"now",
}

// IsKeyword checks if given target string is a valid keyword
func IsKeyword(target string) bool {
	for _, keyword := range keywords {
		if keyword == target {
			return true
		}
	}
	return false
}

// CompileKeyword converts a given string containing a keyword into its associated time.Time
func CompileKeyword(target string) (time.Time, error) {
	for _, keyword := range keywords {
		if keyword == target {
			switch keyword {
			case "now":
				return time.Now(), nil
			}
		}
	}
	return time.Time{}, errors.New("no keyword found")
}
