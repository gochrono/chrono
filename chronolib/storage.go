package chronolib

import (
	"time"
)

// TimespanFilterOptions contains a start and end date to filter frames by
type TimespanFilterOptions struct {
	Start time.Time
	End   time.Time
}

// FrameFilterOptions contains filtering data for retrieving frames. An empty object means no filtering
type FrameFilterOptions struct {
	TimespanFilter TimespanFilterOptions
	Tags           []string
}

// FrameStorage is an interface for a frame storage backend
type FrameStorage interface {
	All(filterOptions FrameFilterOptions) ([]Frame, error)
	Add(frame Frame) (Frame, error)
	Remove(frame Frame) (Frame, error)
	Update(frame Frame) (Frame, error)
}

// StateStorage is an interface for the current frame storage backend
type StateStorage interface {
	Get() (Frame, error)
	Update(frame Frame) (Frame, error)
	Clear() (Frame, error)
}

// GetStateStorage retreives the correct implementation for backend storage
func GetStateStorage() StateStorage {
	statePath := GetAppFilePath("state", "")
	return MsgpackStateFileStorage{statePath}
}

// GetFrameStorage retreives the correct implementation for backend storage
func GetFrameStorage() FrameStorage {
	framesPath := GetAppFilePath("frames", "")
	return MsgpackFrameFileStorage{framesPath}
}
