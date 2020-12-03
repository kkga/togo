package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:     "do [NUM...]",
	Short:   "Mark todo as done",
	Aliases: []string{"d, done"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := "todo.txt"
		var keys []int
		for _, arg := range args {
			key, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Cannot parse todo number: ", arg)
				os.Exit(1)
			}
			keys = append(keys, key)
		}

		m, err := txt.TodoMap(fileName)
		if err != nil {
			fmt.Println("Cannot read todo file:", fileName)
		}

		var completed []txt.Todo
		for _, k := range keys {
			if todo, ok := m[k]; ok {
				todo.ToggleDone()
				if !todo.CreationDate.IsZero() {
					switch todo.Done {
					case true:
						todo.CompletionDate = time.Now()
					case false:
						todo.CompletionDate = time.Time{}
					}
				}
				completed = append(completed, todo)
				m[k] = todo
			} else {
				fmt.Println("Non-existing todo number:", k)
				os.Exit(1)
			}
		}

		if err := txt.WriteTodoMap(m, fileName); err != nil {
			fmt.Println("Cannot write todos to file:", err)
			os.Exit(1)
		}

		if len(completed) > 0 {
			fmt.Println("Marked done:")
			for _, todo := range completed {
				fmt.Println("-", todo.Subject)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
