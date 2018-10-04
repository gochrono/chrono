package chronolib

import (
    "io/ioutil"
    "github.com/vmihailenco/msgpack"
)

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

func SaveState(statePath string, frame *Frame) {
    b, err := msgpack.Marshal(&frame)
    if err != nil {
        panic(err)
    }
    err = ioutil.WriteFile(statePath, b, 0644)
}

func LoadFrames(framesPath string) *Data {
    content, err := ioutil.ReadFile(framesPath)
    var data Data
    if err != nil {
        return &data
    }
    err = msgpack.Unmarshal(content, &data)
    if err != nil {
        panic(err)
    }
    return &data
}

func SaveFrames(framesPath string, data *Data) {
    b, err := msgpack.Marshal(&data)
    if err != nil {
        panic(err)
    }
    err = ioutil.WriteFile(framesPath, b, 0644)
}
