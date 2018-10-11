package chronolib

import (
	"os"
	"os/exec"
)

// GetEditorEnv returns the preferred text editor's binary name
func GetEditorEnv() string {
	return os.Getenv("EDITOR")
}

// EditFile opens a file in an external text editor
func EditFile(editor string, fpath string) {
	editorCmd := exec.Command(editor, fpath)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr
	err := editorCmd.Start()
	if err != nil {
		panic(err)
	}

	err = editorCmd.Wait()
	if err != nil {
		panic(err)
	}

}
