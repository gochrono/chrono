package chronolib

import (
    "strings"
)

// IsAllTags checks that all tags start with a plus
func IsAllTags(tags []string) bool {
    for _, tag := range tags {
        if !strings.HasPrefix(tag, "+") {
            return false
        }
    }
    return true
}


// NormalizeTags tags (trims the prefix, lowercases, and replaces spaces with dashes)
func NormalizeTags(tags []string) []string {
    for idx, tag := range tags {
        tags[idx] = strings.Replace(strings.ToLower(strings.TrimPrefix(tag, "+")), " ", "-", -1)
    }
    return tags
}
