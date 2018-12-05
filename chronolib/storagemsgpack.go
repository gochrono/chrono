package chronolib

import (
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"time"
)

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

// GetAll returns all the frames
func (s MsgpackFrameFileStorage) GetAll() ([]Frame, error) {
	return []Frame{}, nil
}

// GetAllInTimespan gets frames in a specific timespan and filtered based on data in filterOptions
func (s MsgpackFrameFileStorage) GetAllInTimespan(start time.Time, end time.Time, filterOptions FrameFilterOptions) ([]Frame, error) {
	return []Frame{}, nil
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

// GetStateStorage retreives the correct implementation for backend storage
func GetStateStorage() StateStorage {
	statePath := GetAppFilePath("state", "")
	return MsgpackStateFileStorage{statePath}
}
