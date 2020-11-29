package cmd

import "github.com/spf13/cobra"

// RootCmd is the default command
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is CLI task manager",
}
