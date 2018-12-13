package chronolib

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

// NewErrFramesFileDoesNotExist creates a new ErrFileDoesNotExist
func NewErrFramesFileDoesNotExist(message string) *ErrFramesFileDoesNotExist {
	return &ErrFramesFileDoesNotExist{message}
}
