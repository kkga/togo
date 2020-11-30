package cmd

import "github.com/spf13/cobra"

// RootCmd is the default command
var RootCmd = &cobra.Command{
	Use:   "togo [command]",
	Short: "togo is a CLI for todo.txt",
}
