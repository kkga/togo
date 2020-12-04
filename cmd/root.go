package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigFile defines the path to togo.toml configuration file
var ConfigFile string

// TodoFile defines the path to todo.txt
var TodoFile string

var localTodo bool

var rootCmd = &cobra.Command{
	Use:   "togo [command]",
	Short: "togo is a CLI for todo.txt",
}

// Execute executes the rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&ConfigFile, "config", "c", "", "path to config file")
	rootCmd.PersistentFlags().BoolVarP(&localTodo, "local", "l", false, "use todo.txt in current directory")

}

func initConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		viper.AddConfigPath(home + "/.config/togo/")
		viper.SetConfigName("togo")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println(err)
		} else {
			fmt.Println(err)
		}
	}

	switch localTodo {
	case true:
		TodoFile = "todo.txt"
	case false:
		TodoFile = viper.GetString("global.todo_file")
	}
}
