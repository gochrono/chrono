package main

import (
    "errors"
    "fmt"
    "os"
    "github.com/magefile/mage/sh"
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

func Clean() {
    fmt.Println("Cleaning...")
    os.RemoveAll("dist")
}
