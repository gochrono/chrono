package chronolib

import (
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
)

// LoadState returns the current state using the MessagePack format
func LoadState(statePath string) *Frame {
	content, err := ioutil.ReadFile(statePath)
	var frame Frame
	if err != nil {
		return &frame
	}
	err = msgpack.Unmarshal(content, &frame)
	if err != nil {
		panic(err)
	}
	return &frame
}

// SaveState saves the current state using the MessagePack format
func SaveState(statePath string, frame *Frame) {
	b, err := msgpack.Marshal(&frame)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(statePath, b, 0644)
	if err != nil {
		panic(err)
	}
}

// LoadFrames loads frames using the MessagePack format
func LoadFrames(framesPath string) *Data {
	var data Data
	if _, err := os.Stat(framesPath); os.IsNotExist(err) {
		return &Data{}
	}
	content, err := ioutil.ReadFile(framesPath)
	if err != nil {
		panic(err)
	}
	err = msgpack.Unmarshal(content, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

// SaveFrames saves frames using the MessagePack format
func SaveFrames(framesPath string, data *Data) {
	b, err := msgpack.Marshal(&data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(framesPath, b, 0644)
	if err != nil {
		panic(err)
	}
}
