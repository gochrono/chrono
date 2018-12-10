package chronolib

import (
	"time"
)

// ErrFileDoesNotExist represents when a file doesn't exist on the file system
type ErrFileDoesNotExist struct {
	message string
}

// ErrStateFileDoesNotExist represents when a file doesn't exist on the file system
type ErrStateFileDoesNotExist struct {
	message string
}

// ErrFrameNotFound means a frame wasn't found
type ErrFrameNotFound struct {
	message string
}

// Error returns the error message
func (e *ErrFrameNotFound) Error() string {
	return e.message
}

// ErrFramesFileDoesNotExist represents when a file doesn't exist on the file system
type ErrFramesFileDoesNotExist struct {
	message string
}

// Error returns the error message
func (e *ErrFileDoesNotExist) Error() string {
	return e.message
}

// Error returns the error message
func (e *ErrStateFileDoesNotExist) Error() string {
	return e.message
}

// Error returns the error message
func (e *ErrFramesFileDoesNotExist) Error() string {
	return e.message
}

// NewErrFrameNotFound creates a new ErrFrameNotFound
func NewErrFrameNotFound(message string) *ErrFrameNotFound {
	return &ErrFrameNotFound{message}
}

// NewErrFileDoesNotExist creates a new ErrFileDoesNotExist
func NewErrFileDoesNotExist(message string) *ErrFileDoesNotExist {
	return &ErrFileDoesNotExist{message}
}

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

// GetStateStorage retreives the correct implementation for backend storage
func GetStateStorage(config ChronoConfig) StateStorage {
	return MsgpackStateFileStorage{config}
}

// GetFrameStorage retreives the correct implementation for backend storage
func GetFrameStorage(config ChronoConfig) FrameStorage {
	return MsgpackFrameFileStorage{config}
}
