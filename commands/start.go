package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
	"time"
)

func newStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start [project name] [tags]",
		Short: "Start a new frame",
		Long:  "Starts a new frame with the given project name and tags",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			statePath := chronolib.GetAppFilePath("state", "")

			var project = args[0]
			var tags = args[1:]
			var now = time.Now()

			if !chronolib.IsAllTags(tags) {
				fmt.Println("Invalid tag")
				os.Exit(-1)
			}

			tags = chronolib.NormalizeTags(tags)

			frame := chronolib.Frame{
				UUID: []byte{}, Project: project, StartedAt: now, EndedAt: time.Time{}, Tags: tags, Notes: []string{}}

			b, err := msgpack.Marshal(&frame)
			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(statePath, b, 0644)
			if err != nil {
				panic(err)
			}

			fmt.Println(chronolib.FormatNewFrameMessage(frame))
		},
	}
}
