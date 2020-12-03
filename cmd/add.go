package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [TODO]",
	Short:   "Add todo",
	Aliases: []string{"a"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t := strings.Join(args, " ")
		todo, err := txt.AddTodo(t, "todo.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\"\n", todo.Subject)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
