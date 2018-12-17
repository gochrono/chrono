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

// FrameDeleteOptions contains data needed to remove a frame
type FrameDeleteOptions struct {
	Target string
}

// FrameGetOptions contains filtering information needed to find a single Frame
type FrameGetOptions struct {
	Target string
}

// FrameStorage is an interface for a frame storage backend
type FrameStorage interface {
	All(filterOptions FrameFilterOptions) ([]Frame, error)
	Add(frame Frame) (Frame, error)
	Get(getOptions FrameGetOptions) (Frame, error)
	Delete(deleteOptions FrameDeleteOptions) (Frame, error)
	Update(frame Frame) (Frame, error)
	Projects() ([]string, error)
	Tags() ([]string, error)
}

// StateStorage is an interface for the current frame storage backend
type StateStorage interface {
	Get() (Frame, error)
	Update(frame Frame) (Frame, error)
	Clear() (Frame, error)
}
