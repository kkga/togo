package cmd

import (
	"github.com/kkga/task/txt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		txt.AllTasks()
		// tasks, err := txt.AllTasks()
		// if err != nil {
		// 	fmt.Println("Something went wrong: ", err.Error())
		// 	os.Exit(1)
		// }
		// if len(tasks) == 0 {
		// 	fmt.Println("You have no tasks.")
		// 	return
		// }
		// fmt.Println("All tasks:")
		// for i, task := range tasks {
		// 	fmt.Printf("%d. %s\n", i+1, task.Value)
		// }
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
