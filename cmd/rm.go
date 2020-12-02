package cmd

import (
	"fmt"
	"strconv"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove todo",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Cannot parse todo number")
		}

		removedTodo, err := txt.DeleteTodo(key)
		if err != nil {
			fmt.Println("Something went wrong")
		}

		fmt.Println("Removed:", removedTodo.Subject)
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
