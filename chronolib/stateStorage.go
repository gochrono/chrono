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

// StateFilename is the filename of the states file
const StateFilename = "state"

// StateRepo is an interface for retieving current state
type StateRepo interface {
	Load() (State, error)
	Save(state State) error
}

// JSONStateRepo retrieves state in msgpack format
type JSONStateRepo struct {
	config ChronoConfig
}

func writeBytes(config ChronoConfig, filename string, b []byte) error {
	statePath := filepath.Join(config.ConfigDir, filename+"."+config.StorageType)
	jww.INFO.Printf("saving state to %s", statePath)
	err := ioutil.WriteFile(statePath, b, 0644)
	if err != nil {
		jww.INFO.Printf("error writing state")
		return err
	}
	return nil
}

func loadBytes(config ChronoConfig, filename string) ([]byte, error) {
	statePath := filepath.Join(config.ConfigDir, filename+"."+config.StorageType)
	jww.INFO.Printf("reading state from %s", statePath)
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		jww.INFO.Printf("no state found, loading empty state")
		return []byte{}, nil
	}
	b, err := ioutil.ReadFile(statePath)
	if err != nil {
		jww.INFO.Printf("error reading state, loading empty state")
		return []byte{}, nil
	}
	return b, nil
}

// Save writes the state to the current file
func (s JSONStateRepo) Save(state State) error {
	b, err := json.Marshal(state.Get())
	if err != nil {
		jww.INFO.Printf("error marshaling state")
		return err
	}
	return writeBytes(s.config, StateFilename, b)
}

// Load reads the state
func (s JSONStateRepo) Load() (State, error) {
	var currentFrame CurrentFrame
	content, err := loadBytes(s.config, StateFilename)
	if err != nil {
		jww.INFO.Printf("error unmarshling state, loading empty state")
		return State{CurrentFrame{}}, err
	}
	if err = json.Unmarshal(content, &currentFrame); err != nil {
		jww.INFO.Printf("error unmarshling state, loading empty state")
		return State{CurrentFrame{}}, err
	}
	return State{currentFrame}, nil
}

// MsgpackStateRepo retrieves state in msgpack format
type MsgpackStateRepo struct {
	config ChronoConfig
}

// Save writes the state to the current file
func (s MsgpackStateRepo) Save(state State) error {
	b, err := msgpack.Marshal(state.Get())
	if err != nil {
		jww.INFO.Printf("error marshaling state")
		return err
	}
	return writeBytes(s.config, StateFilename, b)
}

// Load reads the state
func (s MsgpackStateRepo) Load() (State, error) {
	var currentFrame CurrentFrame
	content, err := loadBytes(s.config, StateFilename)
	if err != nil {
		jww.INFO.Printf("error unmarshling state, loading empty state")
		return State{CurrentFrame{}}, nil
	}
	if err = msgpack.Unmarshal(content, &currentFrame); err != nil {
		jww.INFO.Printf("error unmarshling state, loading empty state")
		return State{CurrentFrame{}}, nil
	}
	return State{currentFrame}, nil
}

func getStateStorage(config ChronoConfig) StateRepo {
	switch config.StorageType {
	case "json":
		return JSONStateRepo{config}
	case "msgpack":
		return MsgpackStateRepo{config}
	default:
		panic(errors.New("unknown storage type"))
	}
}

// GetState retrieves the state
func GetState(config ChronoConfig) (State, error) {
	stateRepo := getStateStorage(config)
	return stateRepo.Load()
}

// SaveState writes the state to the file
func SaveState(config ChronoConfig, state State) error {
	stateRepo := getStateStorage(config)
	return stateRepo.Save(state)
}
