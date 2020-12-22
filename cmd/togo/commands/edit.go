package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit [NUM]",
	Short:   "Open todo.txt in editor",
	Aliases: []string{"e"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vim"
		}
		edit := exec.Command(editor, TodoFile)
		if len(args) > 0 {
			edit = exec.Command(editor, TodoFile, "+"+args[0])
		}
		edit.Stdin = strings.NewReader("")
		edit.Stdout = os.Stdout
		edit.Stderr = os.Stderr

		err := edit.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
