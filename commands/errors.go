package commands

// ErrTimeStringNotValid represents when a file doesn't exist on the file system
type ErrTimeStringNotValid struct {
	message string
}

// Error returns the error message
func (e *ErrTimeStringNotValid) Error() string {
	return e.message
}

// NewErrTimeStringNotValid return a ErrTimeStringNotValid
func NewErrTimeStringNotValid(message string) *ErrTimeStringNotValid {
	return &ErrTimeStringNotValid{message}
}

// ErrTagNotValid represents when a file doesn't exist on the file system
type ErrTagNotValid struct {
	message string
}

// Error returns the error message
func (e *ErrTagNotValid) Error() string {
	return e.message
}

// NewErrTagNotValid return a ErrTimeStringNotValid
func NewErrTagNotValid(message string) *ErrTagNotValid {
	return &ErrTagNotValid{message}
}
