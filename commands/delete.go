package commands

import (
	"strings"
	"os"
	"bufio"
	"strconv"
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)


func ConfirmDelete(frame chronolib.Frame) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(chronolib.FormatFrameDelete(frame))
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text == "y"
}

func GetFrame(frames chronolib.Frames, target string) (chronolib.Frame, bool) {
	index, err := strconv.Atoi(target)
	if err == nil {
		frame, ok := frames.GetByIndex(index)
		if ok {
			return frame, true
		}
	}
	return frames.GetByUUID(target)
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
			frames, _ := chronolib.GetFrames(config)

			frame, ok := GetFrame(frames, args[0])

			if !ok {
				fmt.Println("Could not find frame")
				return
			}
			shouldDelete := ConfirmDelete(frame)
			if shouldDelete {
				frames.Delete(frame)
				chronolib.SaveFrames(config, frames)
			}
		},
	}
}
