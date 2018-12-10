package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
)

const mainDescription = `Chrono is a time to help track what you spend your time on.

You can start tracking your time with ` + "`start`" + ` and you can
stop the timer with ` + "`stop`" + `.`

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var banner = `   _____ _
  / ____| |
 | |    | |__  _ __ ___  _ __   ___
 | |    | '_ \| '__/ _ \| '_ \ / _ \
 | |____| | | | | | (_) | | | | (_) |
  \_____|_| |_|_|  \___/|_| |_|\___/`

var versionTemplate = fmt.Sprintf(`%s

Version: %s
Commit: %s
Built: %s`, banner, version, commit, date+"\n")

func PrintErrorAndExit(e error) {
	switch e.(type) {
	case *chronolib.ErrFileDoesNotExist:
		fmt.Println("error: " + commandError.Error())
		os.Exit(-1)
	case *chronolib.ErrStateFileDoesNotExist:
		fmt.Println(chronolib.FormatNoProjectMessage())
		os.Exit(-2)
	case *chronolib.ErrFramesFileDoesNotExist:
		fmt.Println(chronolib.FormatNoFramesMessage())
		os.Exit(-3)
	default:
		panic(commandError)
	}
}

var commandError error

// Execute creates the root command with all sub-commands installed
func Execute() {
	var rootCmd = &cobra.Command{
		Use:     "chrono",
		Long:    mainDescription,
		Version: version,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if commandError != nil {
			}
		},
	}
	rootCmd.SetVersionTemplate(versionTemplate)
	rootCmd.AddCommand(newStartCmd(), newStatusCmd(), newStopCmd(), newReportCmd(),
		newLogCmd(), newCancelCmd(), newDeleteCmd(), newFramesCmd(),
		newProjectsCmd(), newRestartCmd(),
		newEditCmd(), newVersionCmd(), newNotesCmd(), newTagsCmd())
	rootCmd.Execute()
}
