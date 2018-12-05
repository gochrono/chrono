package chronolib

import (
	"errors"
	"strings"
)

// CheckTags returns nil if all tags are valid, otherwise it returns a error
func CheckTags(tags []string) error {
	for _, tag := range tags {
		if !strings.HasPrefix(tag, "+") {
			return errors.New("invalid tag: " + tag)
		}
	}
	return nil
}

// NormalizeTags tags (trims the prefix, lowercases, and replaces spaces with dashes)
func NormalizeTags(tags []string) []string {
	for idx, tag := range tags {
		tags[idx] = strings.Replace(strings.ToLower(strings.TrimPrefix(tag, "+")), " ", "-", -1)
	}
	return tags
}
