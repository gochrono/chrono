package chronolib

import (
    "sort"
    "errors"
    "time"
    "github.com/jinzhu/now"
)

// FrameRaw is used for converting to and from JSON
type FrameRaw struct {
    Project string
    StartedAt string
    EndedAt string
    Tags []string
    Notes []string
}

// Frame is a timespan containing some metadata
type Frame struct {
    UUID []byte
    Project string
    StartedAt time.Time
    EndedAt time.Time
    Tags []string
    Notes []string
}

// Data is a wrapper for frames
type Data struct {
    Frames []Frame
}


// GetFrameByIndex retrieves a frame by its index
func (d Data) GetFrameByIndex(index int) (Frame, error) {
    if index <= 0 && index >= len(d.Frames) {
        return Frame{}, errors.New("No frame found")
    }
    return d.Frames[index], nil
}

// GetFrameByShortHex gets a frame using the short form of its UUID
func (d Data) GetFrameByShortHex(hex string) (int, Frame, error) {
    for idx, frame := range d.Frames {
        if GetShortHex(frame.UUID) == hex {
            return idx, frame, nil
        }
    }
    return 0, Frame{}, errors.New("No frame found")
}

// SortFramesByDate sorts frame by its start date
func SortFramesByDate(frames []Frame) {
    sort.Slice(frames, func(i, j int) bool {
        return frames[i].StartedAt.Before(frames[j].StartedAt)
    })
}

// SortTimeMapKeys sorts a timemap by its key (a time.Time)
func SortTimeMapKeys(timemap *map[time.Time][]Frame) []time.Time {
    var keys []time.Time
    for k := range *timemap {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(i, j int) bool {
        return keys[i].Before(keys[j])
    })
    return keys
}

// ConvertFrameToRawFrame converts a frame to a raw frame
func ConvertFrameToRawFrame(frame Frame) FrameRaw {
    return FrameRaw{
        frame.Project,
        frame.StartedAt.Format("2006-01-02 15:04:05"),
        frame.EndedAt.Format("2006-01-02 15:04:05"),
        frame.Tags,
        frame.Notes,
    }
}

// ConvertRawFrameToFrame converts a raw frame back to a frame
func ConvertRawFrameToFrame(uuid []byte, rawFrame FrameRaw) Frame {
    started, err := now.Parse(rawFrame.StartedAt)
    if err != nil { panic(err) }
    ended, err := now.Parse(rawFrame.EndedAt)
    if err != nil { panic(err) }
    return Frame{
        uuid,
        rawFrame.Project,
        started,
        ended,
        rawFrame.Tags,
        rawFrame.Notes,
    }
}

// ContainsMoreThanOneBooleanFlag is a helper method for checking if more than one boolean is true
func ContainsMoreThanOneBooleanFlag(flags... bool) bool {
    count := 0
    for _, flag := range flags {
        if flag {
            count += 1
        }
        if count == 2 {
            return true
        }
    }
    return false
}

