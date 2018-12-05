package chronolib

import (
	"time"
)

// FrameFilterOptions contains filtering data for retrieving frames. An empty object means no filtering
type FrameFilterOptions struct {
	Tags []string
}

// FrameStorage is an interface for a frame storage backend
type FrameStorage interface {
	GetAll() ([]Frame, error)
	GetAllInTimespan(start time.Time, end time.Time, filterOptions FrameFilterOptions) ([]Frame, error)
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
