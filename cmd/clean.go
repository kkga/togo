package cmd

import (
	"fmt"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"cl"},
	Short:   "Move done todos to done.txt",
	Run: func(cmd *cobra.Command, args []string) {
		removed, err := txt.CleanTodos("todo.txt")
		if err != nil {
			fmt.Println(err)
		}

		for _, todo := range removed {
			fmt.Println("Archived:", todo.Subject)
		}

	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}
