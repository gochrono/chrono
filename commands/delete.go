package commands

import (
	"bufio"
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// ConfirmDelete asks user if they are sure they want to delete the frame
func ConfirmDelete(frame chronolib.Frame) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(chronolib.FormatFrameDelete(frame))
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.Replace(text, "\n", "", -1))
	return text == "y"
}

var deleteForce bool

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
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
			var shouldDelete bool
			if deleteForce {
				shouldDelete = true
			} else {
				shouldDelete = ConfirmDelete(frame)
			}
			if shouldDelete {
				frames.Delete(frame)
				chronolib.SaveFrames(config, frames)
			}
		},
	}
	cmd.Flags().BoolVarP(&deleteForce, "force", "f", false, "delete frame without confirmation frame")
	return cmd
}
