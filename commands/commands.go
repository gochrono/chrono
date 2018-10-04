package commands

import (
    "github.com/spf13/cobra"
)

func Execute() {
    var rootCmd = &cobra.Command{Use: "chrono"}
    rootCmd.AddCommand(newStartCmd(), newStatusCmd(), newStopCmd(), newLogCmd(), newEditCmd(), newVersionCmd())
    rootCmd.Execute()
}
