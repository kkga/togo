package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kkga/gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\"\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
