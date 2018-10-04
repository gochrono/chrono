package chronolib

import (
    "sort"
    "errors"
    "time"
    "github.com/jinzhu/now"
)

type FrameRaw struct {
    Project string
    StartedAt string
    EndedAt string
    Tags []string
    Notes []string
}

type Frame struct {
    UUID []byte
    Project string
    StartedAt time.Time
    EndedAt time.Time
    Tags []string
    Notes []string
}

type Data struct {
    Frames []Frame
}


func (d Data) GetFrameByIndex(index int) (Frame, error) {
    if index <= 0 && index >= len(d.Frames) {
        return Frame{}, errors.New("No frame found")
    }
    return d.Frames[index], nil
}

func (d Data) GetFrameByShortHex(hex string) (int, Frame, error) {
    for idx, frame := range d.Frames {
        if GetShortHex(frame.UUID) == hex {
            return idx, frame, nil
        }
    }
    return 0, Frame{}, errors.New("No frame found")
}

func SortFramesByDate(frames []Frame) {
    sort.Slice(frames, func(i, j int) bool {
        return frames[i].StartedAt.Before(frames[j].StartedAt)
    })
}

func SortTimeMapKeys(timemap *map[time.Time][]Frame) []time.Time {
    var keys[] time.Time
    for k := range *timemap {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(i, j int) bool {
        return keys[i].Before(keys[j])
    })
    return keys
}

func ConvertFrameToRawFrame(frame Frame) FrameRaw {
    return FrameRaw{
        frame.Project,
        frame.StartedAt.Format("2006-01-02 15:04:05"),
        frame.EndedAt.Format("2006-01-02 15:04:05"),
        frame.Tags,
        frame.Notes,
    }
}

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

