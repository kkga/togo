package commands

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/kkga/togo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lsCmd = &cobra.Command{
	Use:     "ls [query...]",
	Short:   "List todos",
	Example: "togo ls\ntogo ls +myproject\ntogo ls myquery",
	Aliases: []string{"l", "list"},
	Run: func(cmd *cobra.Command, args []string) {
		m, err := togo.TodoMap(TodoFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		matchingTodos := make(map[int]togo.Todo)
		for k, todo := range m {
			if len(args) > 0 {
				for _, arg := range args {
					_, exists := matchingTodos[k]
					matches := strings.Contains(todo.Subject, arg)
					if !exists && matches {
						matchingTodos[k] = todo
					}
				}
			} else {
				matchingTodos[k] = todo
			}
		}

		// iteration over map happens in random order, so we store the order
		// in a separate slice
		var keys []int
		for k := range matchingTodos {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		var listedKeys []int

		sorting := viper.GetString("sort")
		switch sorting {
		case "file":
			for _, k := range keys {
				PrintTodo(k, matchingTodos[k])
			}
		case "project":
			projects, err := togo.Projects(matchingTodos)
			sort.Strings(projects)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, p := range projects {
				color := color.New(color.Bold).SprintFunc()
				fmt.Println(color(p) + ":")
				for _, k := range keys {
					if containsString(matchingTodos[k].Projects, p) {
						PrintTodo(k, matchingTodos[k])
						listedKeys = append(listedKeys, k)
					}
				}
			}
		case "context":
			contexts, err := togo.Contexts(matchingTodos)
			sort.Strings(contexts)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, c := range contexts {
				color := color.New(color.Bold).SprintFunc()
				fmt.Println(color(c) + ":")
				for _, k := range keys {
					if containsString(matchingTodos[k].Contexts, c) {
						PrintTodo(k, matchingTodos[k])
						listedKeys = append(listedKeys, k)
					}
				}
			}
		case "prio":
			priorities, err := togo.Priorities(matchingTodos)
			sort.Strings(priorities)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, p := range priorities {
				color := color.New(color.Bold).SprintFunc()
				fmt.Println(color("("+p+")") + ":")
				for _, k := range keys {
					if matchingTodos[k].Priority == p {
						PrintTodo(k, matchingTodos[k])
						listedKeys = append(listedKeys, k)
					}
				}
			}
		}

		if sorting != "file" && len(listedKeys) < len(matchingTodos) {
			color := color.New(color.Bold).SprintFunc()
			fmt.Println(color("(" + "no " + sorting + "):"))
			for _, k := range keys {
				if !containsInt(listedKeys, k) {
					PrintTodo(k, matchingTodos[k])
				}
			}
		}
		fmt.Println("------")
		fmt.Printf("%d/%d todos shown (%s)\n", len(matchingTodos), len(m), TodoFile)
	},
}

// PrintTodo colorizes and outputs the given Todo
func PrintTodo(key int, todo togo.Todo) {
	var result string

	if key == 0 {
		result += "-" + " "
	} else {
		result += fmt.Sprintf("%2d ", key)
	}

	if todo.Done {
		color := color.New(color.Bold).SprintFunc()
		result += color("[x]") + " "
	} else {
		color := color.New(color.Reset).SprintFunc()
		result += color("[ ]") + " "
	}

	if todo.Priority != "" {
		switch todo.Priority {
		case "A":
			color := color.New(color.FgRed, color.Bold).SprintFunc()
			result += color("("+todo.Priority+")") + " "
		case "B":
			color := color.New(color.FgYellow, color.Bold).SprintFunc()
			result += color("("+todo.Priority+")") + " "
		case "C":
			color := color.New(color.FgGreen, color.Bold).SprintFunc()
			result += color("("+todo.Priority+")") + " "
		default:
			color := color.New(color.FgGreen, color.Bold).SprintFunc()
			result += color("("+todo.Priority+")") + " "
		}
	}

	if !todo.CompletionDate.IsZero() {
		color := color.New(color.Reset).SprintFunc()
		result += color(todo.CompletionDate.Format("2006-01-02")) + " "
	}

	if !todo.CreationDate.IsZero() {
		color := color.New(color.Reset).SprintFunc()
		result += color(todo.CreationDate.Format("2006-01-02")) + " "
	}

	if todo.Subject != "" {
		color := color.New(color.Reset).SprintFunc()
		result += color(todo.Subject)
	}

	if len(todo.Projects) > 0 {
		color := color.New(color.FgMagenta).SprintFunc()
		for _, p := range todo.Projects {
			result = strings.ReplaceAll(result, p, color(p))
		}
	}

	if len(todo.Contexts) > 0 {
		color := color.New(color.FgCyan).SprintFunc()
		for _, c := range todo.Contexts {
			result = strings.ReplaceAll(result, c, color(c))
		}
	}

	fmt.Println(result)
}

func containsString(source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}

func containsInt(source []int, value int) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().StringP("sort", "s", "file", "sort order, possible values: \"file\", \"project\", \"context\", \"prio\"")
	viper.BindPFlag("sort", lsCmd.Flags().Lookup("sort"))
}
