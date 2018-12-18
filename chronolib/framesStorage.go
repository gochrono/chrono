package chronolib

import (
	"encoding/json"
	"errors"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
	"path/filepath"
)

const framesFilename = "frames"

// FramesRepo is an interface for loading/saving frames
type FramesRepo interface {
	Load() (Frames, error)
	Save(frames *Frames) error
}

// MsgpackFramesRepo retrieves state in msgpack format
type MsgpackFramesRepo struct {
	config ChronoConfig
}

// JSONFramesRepo stores state in the JSON format
type JSONFramesRepo struct {
	config ChronoConfig
}

// Load frames stored in the JSON format
func (s JSONFramesRepo) Load() (Frames, error) {
	var frames Frames
	content, err := loadBytes(s.config, framesFilename)
	if err != nil {
		jww.INFO.Printf("error unmarshling state, loading empty state")
		return Frames{}, err
	}
	if err = json.Unmarshal(content, &frames); err != nil {
		jww.INFO.Printf("error unmarshling state, loading empty state")
		return Frames{}, err
	}
	return Frames{}, nil
}

// Save writes frames using the JSON format
func (s JSONFramesRepo) Save(f *Frames) error {
	b, err := json.Marshal(f)
	jww.INFO.Printf("JSON marshalled frames = %v ", b)
	if err != nil {
		jww.INFO.Printf("error marshaling state")
		return err
	}
	return writeBytes(s.config, framesFilename, b)
}

// Load retrieves frames from the msgpack format
func (s MsgpackFramesRepo) Load() (Frames, error) {
	var frames Frames
	framesPath := filepath.Join(s.config.ConfigDir, framesFilename+"."+s.config.StorageType)
	jww.INFO.Printf("reading frames from %s", framesPath)
	if _, err := os.Stat(framesPath); os.IsNotExist(err) {
		return Frames{}, nil
	}
	content, err := ioutil.ReadFile(filepath.Clean(framesPath))
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
func (s MsgpackFramesRepo) Save(f *Frames) error {
	framesPath := filepath.Join(s.config.ConfigDir, framesFilename+"."+s.config.StorageType)
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

func getFrameStorage(config ChronoConfig) FramesRepo {
	switch config.StorageType {
	case jsonStorageType:
		return JSONFramesRepo{config}
	case msgpackStorageType:
		return MsgpackFramesRepo{config}
	default:
		panic(errors.New("unknown storage type"))
	}
}

// SaveFrames writes frames
func SaveFrames(config ChronoConfig, frames Frames) error {
	framesRepo := getFrameStorage(config)
	return framesRepo.Save(&frames)
}

// GetFrames writes frames
func GetFrames(config ChronoConfig) (Frames, error) {
	framesRepo := getFrameStorage(config)
	return framesRepo.Load()
}
