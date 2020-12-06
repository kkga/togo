package cmd

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
	Run: func(cmd *cobra.Command, args []string) {
		editor := exec.Command("vim", TodoFile, "+"+args[0])
		editor.Stdin = strings.NewReader("")
		editor.Stdout = os.Stdout
		editor.Stderr = os.Stderr

		err := editor.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
