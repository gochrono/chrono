package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
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

// PrintErrorAndExit prints an appropriate message depending on the error type.
// If the error type is unknown, panics
func PrintErrorAndExit(e error) {
	if e != nil {
		switch e.(type) {
		case *chronolib.ErrStateFileDoesNotExist:
			fmt.Println(chronolib.FormatNoProjectMessage())
			os.Exit(0)
		case *chronolib.ErrFramesFileDoesNotExist:
			fmt.Println(chronolib.FormatNoFramesMessage())
			os.Exit(0)
		case *ErrTimeStringNotValid:
			fmt.Println(chronolib.FormatTimeStringNotValid())
			os.Exit(0)
		default:
			panic(commandError)
		}
	}
}

var commandError error

var verbose bool

// Execute creates the root command with all sub-commands installed
func Execute() {
	jww.SetLogThreshold(jww.LevelFatal)
	jww.SetStdoutThreshold(jww.LevelFatal)
	var rootCmd = &cobra.Command{
		Use:     "chrono",
		Long:    mainDescription,
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				jww.SetLogThreshold(jww.LevelInfo)
				jww.SetStdoutThreshold(jww.LevelInfo)
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
