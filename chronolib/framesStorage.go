package chronolib

import (
	jww "github.com/spf13/jwalterweatherman"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FramesStorage is an interface for loading/saving frames
type FramesStorage interface {
	Load() (Frames, error)
	Save(frames Frames) error
}

// MsgpackFramesRepo retrieves state in msgpack format
type MsgpackFramesRepo struct {
	config ChronoConfig
}

// Load retrieves frames from the msgpack format
func (s *MsgpackFramesRepo) Load() (Frames, error) {
	var frames Frames
	framesPath := filepath.Join(s.config.ConfigDir, FramesFilename)
	jww.INFO.Printf("reading frames from %s", framesPath)
	if _, err := os.Stat(framesPath); os.IsNotExist(err) {
		return Frames{}, nil
	}
	content, err := ioutil.ReadFile(framesPath)
	if err != nil {
		return Frames{}, err
	}
	err = msgpack.Unmarshal(content, &frames)
	if err != nil {
		return Frames{}, err
	}
	return frames, nil
}

// Save writes frames using the msgpack format
func (s *MsgpackFramesRepo) Save(f *Frames) error {
	framesPath := filepath.Join(s.config.ConfigDir, FramesFilename)
	jww.INFO.Printf("saving frames [%v] to %s", f, framesPath)
	b, err := msgpack.Marshal(f)
	jww.INFO.Printf("serialzied frames: %v", b)
	if err != nil {
		jww.INFO.Printf("error marshaling frames")
		return err
	}
	err = ioutil.WriteFile(framesPath, b, 0644)
	if err != nil {
		jww.ERROR.Printf("error writing frames to %s", framesPath)
		return err
	}
	return nil
}

// SaveFrames writes frames
func SaveFrames(config ChronoConfig, frames Frames) error {
	framesRepo := MsgpackFramesRepo{config}
	return framesRepo.Save(&frames)
}

// GetFrames writes frames
func GetFrames(config ChronoConfig) (Frames, error) {
	framesRepo := MsgpackFramesRepo{config}
	return framesRepo.Load()
}
