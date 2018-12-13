package chronolib

import (
	jww "github.com/spf13/jwalterweatherman"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
	"path/filepath"
)

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
		jww.INFO.Printf("error marshaling state")
		return err
	}
	statePath := filepath.Join(s.config.ConfigDir, StateFilename)
	jww.INFO.Printf("saving state to %s", statePath)
	err = ioutil.WriteFile(statePath, b, 0644)
	if err != nil {
		jww.INFO.Printf("error writing state")
		return err
	}
	return nil
}

// Load reads the state
func (s *MsgpackStateRepo) Load() (State, error) {
	statePath := filepath.Join(s.config.ConfigDir, StateFilename)
	jww.INFO.Printf("reading state from %s", statePath)
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		jww.INFO.Printf("no state found, loading empty state")
		return State{CurrentFrame{}}, &ErrStateFileDoesNotExist{statePath}
	}
	content, err := ioutil.ReadFile(statePath)
	var currentFrame CurrentFrame
	if err != nil {
		jww.INFO.Printf("error reading state, loading empty state")
		return State{CurrentFrame{}}, err
	}
	err = msgpack.Unmarshal(content, &currentFrame)
	if err != nil {
		jww.INFO.Printf("error unmarshling state, loading empty state")
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
