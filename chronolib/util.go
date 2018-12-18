package chronolib

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

// StringInSlice checks if a string is present in a slice
func StringInSlice(target string, list []string) bool {
	for _, candidate := range list {
		if candidate == target {
			return true
		}
	}
	return false
}

// FramesEqual check if frames are equal
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

// GetShortHex returns the first six characters from a hex encoded SHA
func GetShortHex(sha string) string {
	return sha[0:7]
}
