package chronolib

import (
	"github.com/satori/go.uuid"
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
	CurrentFrame CurrentFrame
}

// Get retreives the CurrentFrame from the State
func (s *State) Get() CurrentFrame {
	return s.CurrentFrame
}

// Update the CurrentFrame
func (s *State) Update(CurrentFrame CurrentFrame) {
	s.CurrentFrame = CurrentFrame
}

// Clear the CurrentFrame from the state
func (s *State) Clear() {
	s.CurrentFrame = CurrentFrame{}
}

// IsEmpty checks if the CurrentFrame is empty
func (s *State) IsEmpty() bool {
	return s.CurrentFrame.Project == ""
}

// ToFrame converts the CurrentFrame to a Frame by adding a UUID and end time
func (s *State) ToFrame(end time.Time) Frame {
	id := uuid.NewV4()
	return Frame{
		UUID:      id.Bytes(),
		Project:   s.CurrentFrame.Project,
		Tags:      s.CurrentFrame.Tags,
		StartedAt: s.CurrentFrame.StartedAt,
		UpdatedAt: s.CurrentFrame.UpdatedAt,
		EndedAt:   end,
		Notes:     s.CurrentFrame.Notes,
	}
}
