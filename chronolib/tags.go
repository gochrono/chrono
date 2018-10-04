package chronolib

import (
    "strings"
)

func IsAllTags(tags []string) bool {
    for _, tag := range tags {
        if !strings.HasPrefix(tag, "+") {
            return false
        }
    }
    return true
}


func NormalizeTags(tags []string) []string {
    for idx, tag := range tags {
        tags[idx] = strings.Replace(strings.ToLower(strings.TrimPrefix(tag, "+")), " ", "-", -1)
    }
    return tags
}
