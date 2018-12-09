package commands

import (
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
    "github.com/satori/go.uuid"
    "fmt"
)

type FrameV2 struct {
    UUID      string
}


func NewFrameV2() FrameV2 {
    uuid := uuid.Must(uuid.NewV4())
    return FrameV2{uuid.String()}
}

var frames = []FrameV2{
    FrameV2{"f1fb9554-892e-4805-aa90-7ec0e639bfa4"},
    FrameV2{"32e1c5ea-9879-4b66-9950-3df661da0b96"},
    NewFrameV2(),
    NewFrameV2(),
    NewFrameV2(),
    NewFrameV2(),
}


func newDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Deletes a frame through either its index possition or UUID",
		Long:  "Deletes a frame through either its index possition or UUID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
            configDir := chronolib.GetCorrectConfigDirectory("")
            config := chronolib.GetConfig(configDir)
            frameStorage := chronolib.GetFrameStorage(config)
            frame, err := frameStorage.Delete(chronolib.FrameDeleteOptions{
                Target: args[0],
            })
            if err != nil {
                panic(err)
            }
            fmt.Println(frame)
		},
	}
}
