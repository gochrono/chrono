package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func newNotesCmd() *cobra.Command {
	notes := &cobra.Command{
		Use:   "notes",
		Short: "Manage notes for the current frame",
		Long: `Manage notes for the current frame. You can add new notes with the 'add' command
		or remove them with the 'delete' command. You can list all added notes with the
		'show' command.`,
	}
	notes.AddCommand(newNotesAddCmd(), newNotesShowCmd(), newNotesDeleteCmd())
	return notes
}

func newNotesAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [message]",
		Short: "Add a new note to current frame",
		Long:  "Add a new note to current frame",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			stateStorage := chronolib.GetStateStorage()
			state, err := stateStorage.Get()
			if err != nil {
				commandError = err
				return
			}
			state.Notes = append(state.Notes, args[0])
			state, err = stateStorage.Update(state)
			if err != nil {
				commandError = err
				return
			}
		},
	}
}
func newNotesShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show all notes for the current frame",
		Long:  "Show all notes for the current frame",
		Run: func(cmd *cobra.Command, args []string) {
			stateStorage := chronolib.GetStateStorage()
			state, err := stateStorage.Get()
			if err != nil {
				commandError = err
				return
			}
			if len(state.Notes) != 0 {
				for index, note := range state.Notes {
					fmt.Println(chronolib.FormatNoteShowLine(index, note))
				}
			} else {
				fmt.Println(chronolib.FormatNoNotesMessage())
			}
		},
	}
}

func newNotesDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a note from the current frame",
		Long:  "Delete a note from the current frame",
		Run: func(cmd *cobra.Command, args []string) {
			stateStorage := chronolib.GetStateStorage()
			state, err := stateStorage.Get()
			if err != nil {
				commandError = err
				return
			}
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Index must be a number!")
				os.Exit(0)
			} else if index > len(state.Notes) || index < 0 {
				fmt.Println("Index must be a number!")
				os.Exit(0)
			}
			fmt.Printf("Deleting note '%s'\n", state.Notes[index])
			state.Notes = append(state.Notes[:index], state.Notes[index+1:]...)
			state, err = stateStorage.Update(state)
			if err != nil {
				commandError = err
				return
			}
		},
	}
}
