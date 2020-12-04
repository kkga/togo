package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:     "rm [NUM...]",
	Short:   "Remove todo",
	Aliases: []string{"remove"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var keys []int
		for _, arg := range args {
			key, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Cannot parse todo number:", arg)
				os.Exit(1)
			}
			keys = append(keys, key)
		}

		m, err := txt.TodoMap(TodoFile)
		if err != nil {
			fmt.Println("Cannot read todo file:", err)
		}

		var removed []txt.Todo
		for _, k := range keys {
			if todo, ok := m[k]; ok {
				removed = append(removed, todo)
				delete(m, k)
			} else {
				fmt.Println("Cannot find todo number:", k)
			}
		}

		if err := txt.WriteTodoMap(m, TodoFile); err != nil {
			fmt.Println("Cannot write todos to file:", err)
			os.Exit(1)
		}

		if len(removed) > 0 {
			fmt.Println("Removed:")
			for _, todo := range removed {
				fmt.Println("-", todo.Subject)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
