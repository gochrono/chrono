package commands

import (
    "github.com/spf13/cobra"
)

const mainDescription = `Chrono is a time to help track what you spend your time on.

You can start tracking your time with ` + "`start`" + ` and you can
stop the timer with ` + "`stop`" + `.`

func Execute() {
    var rootCmd = &cobra.Command{Use: "chrono", Long: mainDescription}
    rootCmd.AddCommand(newStartCmd(), newStatusCmd(), newStopCmd(),
        newLogCmd(), newEditCmd(), newVersionCmd(), newNotesCmd(), newTagsCmd())
    rootCmd.Execute()
}
