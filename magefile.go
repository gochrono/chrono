// +build mage

package main

import (
    "errors"
    "fmt"
    "os"
    "github.com/magefile/mage/sh"
    "github.com/jordanknott/chrono/chronolib"
    "github.com/google/gofuzz"
)

var Default = Test
var goexe = "go"

func init() {
    if exe := os.Getenv("GOEXE"); exe != "" {
        goexe = exe
    }
}

func Test() error {
    s, err := sh.Output(goexe, "test", "./...")
    if err != nil {
        fmt.Println(s)
        return errors.New("Tests failed")
    }
    fmt.Println("Tests passed")
    return nil
}

func Generate() error {
    f := fuzz.New()
    var frames []chronolib.Frame
    for i := 1; i <= 10000; i++ {
        var frame chronolib.Frame
        f.Fuzz(&frame)
        frames = append(frames, frame)
    }
    chronolib.SaveFrames("testFrames", &chronolib.Data{frames})
    return nil
}

func Clean() {
    fmt.Println("Cleaning...")
    os.RemoveAll("dist")
}
