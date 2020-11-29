package cmd

import (
	"fmt"
	"strconv"

	"github.com/kkga/gophercises/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks tasks as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse: ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			value := task.Value
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as complete\n", id)
			} else {
				fmt.Printf("Marked \"%d: %s\" as complete\n", id, value)
			}
		}

		// for _, id := range ids {
		// 	err := db.DeleteTask(id)
		// 	if err != nil {
		// 		fmt.Println("Failed to delete", err.Error())
		// 	}
		// }
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
