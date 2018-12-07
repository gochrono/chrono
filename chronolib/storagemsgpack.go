package chronolib

import (
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
)

// ErrFileDoesNotExist represents when a file doesn't exist on the file system
type ErrFileDoesNotExist struct {
	message string
}

// Error returns the error message
func (e *ErrFileDoesNotExist) Error() string {
	return e.message
}

// NewErrFileDoesNotExist creates a new ErrFileDoesNotExist
func NewErrFileDoesNotExist(message string) *ErrFileDoesNotExist {
	return &ErrFileDoesNotExist{message}
}

// MsgpackStateFileStorage stores the current frame in the Msgpack format
type MsgpackStateFileStorage struct {
	StatePath string
}

// MsgpackFrameFileStorage stores frames in the Msgpack format
type MsgpackFrameFileStorage struct {
	FramesPath string
}

// Get retrieves the current frame if it exists
func (s MsgpackStateFileStorage) Get() (Frame, error) {
	if _, err := os.Stat(s.StatePath); os.IsNotExist(err) {
		return Frame{}, NewErrFileDoesNotExist(s.StatePath + " does not exist")
	}
	content, err := ioutil.ReadFile(s.StatePath)
	var frame Frame
	if err != nil {
		return Frame{}, err
	}
	err = msgpack.Unmarshal(content, &frame)
	if err != nil {
		return Frame{}, err
	}
	return frame, nil
}

// Update the current frame's information if it exists
func (s MsgpackStateFileStorage) Update(frame Frame) (Frame, error) {
	b, err := msgpack.Marshal(&frame)
	if err != nil {
		return Frame{}, err
	}
	err = ioutil.WriteFile(s.StatePath, b, 0644)
	if err != nil {
		return Frame{}, err
	}
	return frame, nil
}

// Clear the current frame
func (s MsgpackStateFileStorage) Clear() (Frame, error) {
	return Frame{}, nil
}

func getFrames(framesPath string) ([]Frame, error) {
	var data Data
	if _, err := os.Stat(framesPath); os.IsNotExist(err) {
		return []Frame{}, NewErrFileDoesNotExist(framesPath + " does not exist")
	}
	content, err := ioutil.ReadFile(framesPath)
	if err != nil {
		return []Frame{}, err
	}
	err = msgpack.Unmarshal(content, &data)
	if err != nil {
		return []Frame{}, err
	}
	return data.Frames, nil
}

// All returns all frames, filtered using the given FrameFilterOptions
func (s MsgpackFrameFileStorage) All(filterOptions FrameFilterOptions) ([]Frame, error) {
	frames, err := getFrames(s.FramesPath)
	if err != nil {
		return []Frame{}, err
	}
	return FilterFrames(&frames, filterOptions), nil
}

// Add a new frame to storage
func (s MsgpackFrameFileStorage) Add(frame Frame) (Frame, error) {
	return Frame{}, nil
}

// Remove a frame (matched by frame's UUID)
func (s MsgpackFrameFileStorage) Remove(frame Frame) (Frame, error) {
	return Frame{}, nil
}

// Update the information for the given frame (matched by frame's UUID)
func (s MsgpackFrameFileStorage) Update(frame Frame) (Frame, error) {
	return Frame{}, nil
}
