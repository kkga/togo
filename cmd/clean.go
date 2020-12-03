package cmd

import (
	"fmt"
	"os"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"cl"},
	Short:   "Move done todos to done.txt",
	Args:    cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := "todo.txt"
		doneFileName := "done.txt"

		m, err := txt.TodoMap(fileName)
		if err != nil {
			fmt.Println("Cannot read todo file", err)
		}

		var completed []txt.Todo
		for k, todo := range m {
			if todo.Done {
				completed = append(completed, todo)
				delete(m, k)
			}
		}

		if err := txt.WriteTodoMap(m, fileName); err != nil {
			fmt.Println("Cannot write todos to file:", err)
		}

		if len(completed) > 0 {
			fmt.Println("Archived:")
			for _, todo := range completed {
				fmt.Println("-", txt.FormatTodo(todo))
			}
		} else {
			fmt.Println("No completed todos")
		}

		doneFile, err := os.OpenFile(doneFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Cannot open done.txt:", err)
			os.Exit(1)
		}
		defer doneFile.Close()

		for _, todo := range completed {
			todoStr := txt.FormatTodo(todo)
			if _, err := doneFile.WriteString(todoStr + "\n"); err != nil {
				fmt.Println("Cannot write to done.txt:", err)
				os.Exit(1)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}
