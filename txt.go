package togo

import (
	"bufio"
	"os"

	"sort"
)

const dateLayout = "2006-01-02"

// LinesInFile scans the given fileName and returns a slice of strings for each line
func LinesInFile(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	result := []string{}

	for s.Scan() {
		result = append(result, s.Text())
	}

	return result, nil
}

// TodoMap converts a todo.txt file into a map of Todos
func TodoMap(fileName string) (map[int]Todo, error) {
	m := make(map[int]Todo)
	todoLines, err := LinesInFile(fileName)
	if err != nil {
		return nil, err
	}
	for i, line := range todoLines {
		m[i+1] = ParseTodo(line)
	}

	return m, nil
}

// Projects returns a slice of all project strings in a given map of Todos
// TODO turn TodoMap into a type with methods for this and contexts
func Projects(m map[int]Todo) (projects []string, err error) {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		for _, p := range m[k].Projects {
			if contains(projects, p) {
				continue
			}
			projects = append(projects, p)
		}
	}
	return
}

// Contexts returns a slice of all context strings in a given map of Todos
func Contexts(m map[int]Todo) (contexts []string, err error) {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		for _, p := range m[k].Contexts {
			if contains(contexts, p) {
				continue
			}
			contexts = append(contexts, p)
		}
	}
	return
}

// Priorities returns a slice of all priority strings in a given map of Todos
func Priorities(m map[int]Todo) (priorities []string, err error) {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		p := m[k].Priority
		if p != "" && !contains(priorities, p) {
			priorities = append(priorities, p)
		}
	}
	return
}

// WriteTodoMap writes a map of Todos into a formatted file
func WriteTodoMap(m map[int]Todo, fileName string) error {
	todos := make([]string, 0)

	// maps are iterated in random order, so we store ordered keys separately
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		todo := FormatTodo(m[k])
		todos = append(todos, todo)
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	datawriter := bufio.NewWriter(f)
	for _, todo := range todos {
		_, _ = datawriter.WriteString(todo + "\n")
	}
	if err := datawriter.Flush(); err != nil {
		return err
	}

	return nil
}

func contains(source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}
