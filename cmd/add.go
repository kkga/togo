package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [TODO]",
	Short:   "Add todo",
	Aliases: []string{"a"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todoStr := strings.Join(args, " ")

		fileName := "todo.txt"
		m, err := txt.TodoMap(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		todo := txt.ParseTodo(todoStr)
		todo.CreationDate = time.Now()
		m[len(m)+1] = todo

		if err := txt.WriteTodoMap(m, fileName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Added:", todo.Subject)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
