package chronolib

import (
	"github.com/satori/go.uuid"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// CurrentFrame contains data for the current frame
type CurrentFrame struct {
	Project   string
	StartedAt time.Time
	UpdatedAt time.Time
	Tags      []string
	Notes     []string
}

// State contains the CurrentFrame and provides methods for interacting with it
type State struct {
	currentFrame CurrentFrame
}

// Frames contains recordered frames
type Frames struct {
	frames []Frame
}

// StateRepo is an interface for retieving current state
type StateRepo interface {
	Load() (State, error)
	Save(state State) error
}

// MsgpackStateRepo retrieves state in msgpack format
type MsgpackStateRepo struct {
	config ChronoConfig
}

// Save writes the state to the current file
func (s *MsgpackStateRepo) Save(state State) error {
	b, err := msgpack.Marshal(state.Get())
	if err != nil {
		return err
	}
	statePath := filepath.Join(s.config.ConfigDir, StateFilename)
	err = ioutil.WriteFile(statePath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Load reads the state
func (s *MsgpackStateRepo) Load() (State, error) {
	statePath := filepath.Join(s.config.ConfigDir, StateFilename)
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		return State{CurrentFrame{}}, &ErrStateFileDoesNotExist{statePath}
	}
	content, err := ioutil.ReadFile(statePath)
	var currentFrame CurrentFrame
	if err != nil {
		return State{CurrentFrame{}}, err
	}
	err = msgpack.Unmarshal(content, &currentFrame)
	if err != nil {
		return State{CurrentFrame{}}, err
	}
	return State{currentFrame}, nil
}

// GetState retrieves the state
func GetState(config ChronoConfig) (State, error) {
	stateRepo := MsgpackStateRepo{config}
	return stateRepo.Load()
}

// SaveState writes the state to the file
func SaveState(config ChronoConfig, state State) error {
	stateRepo := MsgpackStateRepo{config}
	return stateRepo.Save(state)
}

// Get retreives the CurrentFrame from the State
func (s *State) Get() CurrentFrame {
	return s.currentFrame
}

// Update the CurrentFrame
func (s *State) Update(currentFrame CurrentFrame) {
	s.currentFrame = currentFrame
}

// Clear the CurrentFrame from the state
func (s *State) Clear() {
	s.currentFrame = CurrentFrame{}
}

// IsEmpty checks if the CurrentFrame is empty
func (s *State) IsEmpty() bool {
	return s.currentFrame.Project == ""
}

// ToFrame converts the CurrentFrame to a Frame by adding a UUID and end time
func (s *State) ToFrame(end time.Time) Frame {
	id := uuid.NewV4()
	return Frame{
		UUID:      id.Bytes(),
		Project:   s.currentFrame.Project,
		Tags:      s.currentFrame.Tags,
		StartedAt: s.currentFrame.StartedAt,
		UpdatedAt: s.currentFrame.UpdatedAt,
		EndedAt:   end,
		Notes:     s.currentFrame.Notes,
	}
}

// FramesStorage is an interface for loading/saving frames
type FramesStorage interface {
	Load() (Frames, error)
	Save(frames Frames) error
}

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
