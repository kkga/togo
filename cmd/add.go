package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:     "add [TODO]",
	Short:   "Add todo",
	Aliases: []string{"a"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todoStr := strings.Join(args, " ")

		m, err := txt.TodoMap(TodoFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		addDate := viper.GetBool("prepend_date")
		todo := txt.ParseTodo(todoStr)
		if addDate {
			todo.CreationDate = time.Now()
		}
		m[len(m)+1] = todo

		if err := txt.WriteTodoMap(m, TodoFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Added (%s):\n", TodoFile)
		PrintTodo(0, todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
