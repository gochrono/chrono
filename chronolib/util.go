package chronolib

import (
    "crypto/sha1"
    "encoding/hex"
    "time"
)


func isSlicesEqual(s1, s2 []string) bool {
    if (s1 == nil) != (s2 == nil) {
        return false
    }

    if len(s1) != len(s2) {
        return false
    }

    for index := range s1 {
        if s1[index] != s2[index] {
            return false
        }
    }
    return true
}

func FramesEqual(f1 Frame, f2 Frame) bool {
    if f1.Project != f2.Project {
        return false
    }
    if f1.StartedAt.Format("2006-01-02 15:04:05") != f2.StartedAt.Format("2006-01-02 15:04:05") {
        return false
    }
    if f1.EndedAt.Format("2006-01-02 15:04:05") != f2.EndedAt.Format("2006-01-02 15:04:05") {
        return false
    }
    if isSlicesEqual(f1.Tags, f2.Tags) != true {
        return false
    }
    if isSlicesEqual(f1.Notes, f2.Notes) != true {
        return false
    }
    return true
}

func GetShortHex(sha []byte) string {
    hex := hex.EncodeToString(sha)
    return hex[0:6]
}


func CreateFrameUUID(name string, start *time.Time, end *time.Time) []byte {
    input := []byte(name + start.Format("2006-01-02 15:04:05") + end.Format("2006-01-02 15:04:05"))
    hasher := sha1.New()
    hasher.Write(input)
    return hasher.Sum(nil)
}

