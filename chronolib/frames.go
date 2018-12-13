package chronolib

import (
	jww "github.com/spf13/jwalterweatherman"
)

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
	return Frame{}, false
}

// GetByUUID retrieves a frame by its uuid
func (s *Frames) GetByUUID(id string) (Frame, bool) {
	return Frame{}, false
}

// Add a new frame to the frames list
func (s *Frames) Add(frame Frame) {
	jww.INFO.Printf("adding frame %v to frame list: %v", frame, s.Frames)
	s.Frames = append(s.Frames, frame)
}

// Update a frame in the frames list, matched by its UUID
func (s *Frames) Update(frame Frame) {
}

// Remove a frame from the frames list, matched by its UUID
func (s *Frames) Remove(frame Frame) {
}
