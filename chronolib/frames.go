package chronolib

import (
	jww "github.com/spf13/jwalterweatherman"
	"time"
)

// Frame is a timespan containing some metadata
type Frame struct {
	UUID      string
	Project   string
	StartedAt time.Time
	EndedAt   time.Time
	UpdatedAt time.Time
	Tags      []string
	Notes     []string
}

// Frames contains recordered frames
type Frames struct {
	Frames []Frame
}

// All retrieves all frames
func (s *Frames) All() []Frame {
	return s.Frames
}

// Filter retrieves frames based on filter options
func (s *Frames) Filter(filterOptions FrameFilterOptions) []Frame {
	return FilterFrames(&s.Frames, filterOptions)
}

// GetByIndex retrieves a frame by its index
func (s *Frames) GetByIndex(index int) (Frame, bool) {
	var targetIndex int
	if index >= 0 {
		targetIndex = index
	} else {
		targetIndex = index + len(s.Frames)
	}
	if targetIndex >= len(s.Frames) || targetIndex < 0 {
		return Frame{}, false
	}
	return s.Frames[targetIndex], true
}

// GetByUUID retrieves a frame by its uuid
func (s *Frames) GetByUUID(id string) (Frame, bool) {
	for _, frame := range s.Frames {
		if GetShortHex(frame.UUID) == id {
			return frame, true
		}
	}
	return Frame{}, false
}

// Add a new frame to the frames list
func (s *Frames) Add(frame Frame) {
	jww.INFO.Printf("adding frame %v to frame list: %v", frame, s.Frames)
	s.Frames = append(s.Frames, frame)
}

func (s *Frames) GetFrameIndex(target Frame) int {
	for index, frame := range s.Frames {
		if frame.UUID == target.UUID {
			return index
		}
	}
	jww.ERROR.Printf("can not find frame by index: %v", target)
	return 0
}

// Update a frame in the frames list, matched by its UUID
func (s *Frames) Update(target Frame) {
	index := s.GetFrameIndex(target)
	s.Frames[index] = target
}

// Delete a frame from the frames list, matched by its UUID
func (s *Frames) Delete(targetFrame Frame) {
	for index, frame := range s.Frames {
		if GetShortHex(frame.UUID) == GetShortHex(targetFrame.UUID) {
			jww.INFO.Printf("deleting frame %v", s.Frames[index])
			s.Frames = append(s.Frames[:index], s.Frames[index+1:]...)
			break
		}
	}
}

// Tags returns a list of all unique tags found in all frames
func (s *Frames) Tags() []string {
	encountered := map[string]bool{}
	for _, frame := range s.Frames {
		for _, tag := range frame.Tags {
			encountered[tag] = true
		}
	}
	keys := make([]string, 0, len(encountered))
	for k := range encountered {
		keys = append(keys, k)
	}
	return keys
}

// Projects returns a list of all unique project names found in all frames
func (s *Frames) Projects() []string {
	encountered := map[string]bool{}
	for _, frame := range s.Frames {
		encountered[frame.Project] = true
	}
	keys := make([]string, 0, len(encountered))
	for k := range encountered {
		keys = append(keys, k)
	}
	return keys
}
