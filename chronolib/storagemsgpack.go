package chronolib

import (
	"github.com/vmihailenco/msgpack"
	"path/filepath"
	"io/ioutil"
	"os"
    "strings"
	"encoding/hex"
    "strconv"
)

const StateFilename = "state.msgpack"
const FramesFilename = "frames.msgpack"

// MsgpackStateFileStorage stores the current frame in the Msgpack format
type MsgpackStateFileStorage struct {
	Config ChronoConfig
}

// MsgpackFrameFileStorage stores frames in the Msgpack format
type MsgpackFrameFileStorage struct {
	Config ChronoConfig
}

func (s *MsgpackStateFileStorage) GetPath() string {
	return filepath.Join(s.Config.ConfigDir, FramesFilename)
}

func (s *MsgpackFrameFileStorage) GetPath() string {
	return filepath.Join(s.Config.ConfigDir, StateFilename)
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

func saveFrames(framesPath string, frames []Frame) error {
	b, err := msgpack.Marshal(Data{frames})
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(framesPath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}


func saveState(statePath string, frame Frame) (Frame, error) {
	b, err := msgpack.Marshal(&frame)
	if err != nil {
		return Frame{}, err
	}
	err = ioutil.WriteFile(statePath, b, 0644)
	if err != nil {
		return Frame{}, err
	}
	return frame, nil
}

// Get retrieves the current frame if it exists
func (s MsgpackStateFileStorage) Get() (Frame, error) {
    if _, err := os.Stat(s.GetPath()); os.IsNotExist(err) {
        return Frame{}, &ErrStateFileDoesNotExist{s.GetPath()}
	}
    content, err := ioutil.ReadFile(s.GetPath())
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
    return saveState(s.GetPath(), frame)
}

// Clear the current frame
func (s MsgpackStateFileStorage) Clear() (Frame, error) {
    return saveState(s.GetPath(), Frame{})
}

// All returns all frames, filtered using the given FrameFilterOptions
func (s MsgpackFrameFileStorage) All(filterOptions FrameFilterOptions) ([]Frame, error) {
    frames, err := getFrames(s.GetPath())
	if err != nil {
		return []Frame{}, err
	}
	return FilterFrames(&frames, filterOptions), nil
}

// Add a new frame to storage
func (s MsgpackFrameFileStorage) Add(frame Frame) (Frame, error) {
    frames, err := getFrames(s.GetPath())
	if err != nil {
		switch err.(type) {
		case *ErrFileDoesNotExist:
			frames = []Frame{}
		default:
			return Frame{}, err
		}
	}
	frames = append(frames, frame)
    err = saveFrames(s.GetPath(), frames)
	if err != nil {
		return Frame{}, err
	}
	return frame, nil
}

// Projects returns a unique list of all project names used in frames
func (s MsgpackFrameFileStorage) Projects() ([]string, error) {
    frames, err := getFrames(s.GetPath())
	if err != nil {
		return []string{}, err
	}
	encountered := map[string]bool{}
	for _, frame := range frames {
        encountered[frame.Project] = true
	}
	keys := make([]string, 0, len(encountered))
	for k := range encountered {
		keys = append(keys, k)
	}
	return keys, nil
}

// Tags returns a unique list of all tags used in frames
func (s MsgpackFrameFileStorage) Tags() ([]string, error) {
    frames, err := getFrames(s.GetPath())
	if err != nil {
		return []string{}, err
	}
	encountered := map[string]bool{}
	for _, frame := range frames {
		for _, tag := range frame.Tags {
			encountered[tag] = true
		}
	}
	keys := make([]string, 0, len(encountered))
	for k := range encountered {
		keys = append(keys, k)
	}
	return keys, nil
}

// Delete a frame (matched by frame's UUID)
func (s MsgpackFrameFileStorage) Delete(deleteOptions FrameDeleteOptions) (Frame, error) {
    frames, err := s.All(FrameFilterOptions{})
    if err != nil {
        return Frame{}, err
    }
    idx, err := strconv.Atoi(deleteOptions.Target)
    if err == nil {
        var targetIndex int
        if idx >= 0 {
            targetIndex = idx
        } else {
            targetIndex = idx + len(frames)
        }
        
        if targetIndex >= 0 && targetIndex < len(frames) {
	        frames = append(frames[:targetIndex], frames[targetIndex+1:]...)
            return frames[targetIndex], nil
        }
    }
    for targetIndex, frame := range frames {
        if strings.HasPrefix(hex.EncodeToString(frame.UUID), deleteOptions.Target) {
	        frames = append(frames[:targetIndex], frames[targetIndex+1:]...)
            err = saveFrames(s.GetPath(), frames)
            return frame, err
        }
    }
    return Frame{}, NewErrFrameNotFound(deleteOptions.Target)
}

// Update the information for the given frame (matched by frame's UUID)
func (s MsgpackFrameFileStorage) Update(newFrame Frame) (Frame, error) {
    frames, err := s.All(FrameFilterOptions{})
    if err != nil {
        return Frame{}, err
    }
    for index, frame := range frames {
        if hex.EncodeToString(frame.UUID) == hex.EncodeToString(newFrame.UUID) {
            frames[index] = newFrame
            return newFrame, nil
        }
    }
	return newFrame, nil
}
