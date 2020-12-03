package txt

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

// Todo represents a single todo item
type Todo struct {
	Done           bool
	Priority       string
	CompletionDate time.Time
	CreationDate   time.Time
	Subject        string
	Project        string
	Context        string
	Tags           []string
}

// ToggleDone marks the todo as complete or incomplete
func (t *Todo) ToggleDone() bool {
	t.Done = !t.Done
	return t.Done
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

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

// ParseTodo converts a string into a Todo struct
func ParseTodo(todo string) Todo {
	var done bool
	var subject string
	var completionDate time.Time
	var creationDate time.Time

	// check if the task is done
	if strings.HasPrefix(todo, "x ") {
		done = true
		todo = strings.Replace(todo, "x ", "", 1)
	}

	// parse dates
	dateRe := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	dates := dateRe.FindAllString(todo, -1)

	if len(dates) == 1 {
		todo = strings.Replace(todo, dates[0]+" ", "", 1)
		date, _ := time.Parse(dateLayout, dates[0])
		creationDate = date
	} else if len(dates) > 1 {
		todo = strings.Replace(todo, dates[0]+" ", "", 1)
		compDate, _ := time.Parse(dateLayout, dates[0])
		completionDate = compDate
		todo = strings.Replace(todo, dates[1]+" ", "", 1)
		creatDate, _ := time.Parse(dateLayout, dates[1])
		creationDate = creatDate
	}

	subject = todo

	return Todo{
		Done:           done,
		Subject:        subject,
		CompletionDate: completionDate,
		CreationDate:   creationDate,
	}
}

// FormatTodo converts a Todo struct into a formatted string for output
func FormatTodo(todo Todo) string {
	s := make([]string, 0)

	if todo.Done {
		s = append(s, "x")
	}
	if !todo.CompletionDate.IsZero() {
		s = append(s, todo.CompletionDate.Format(dateLayout))
	}
	if !todo.CreationDate.IsZero() {
		s = append(s, todo.CreationDate.Format(dateLayout))
	}
	if todo.Subject != "" {
		s = append(s, todo.Subject)
	}

	return strings.Join(s, " ")
}

// ListTodos returns a map of formatted todo strings that match given queries
func ListTodos(queries []string, fileName string) (map[int]string, error) {
	m, err := TodoMap(fileName)
	if err != nil {
		return nil, err
	}

	todos := make(map[int]string)

	for k, todo := range m {
		if len(queries) > 0 {
			for _, q := range queries {
				_, exists := todos[k]
				matches := strings.Contains(todo.Subject, q)
				if !exists && matches {
					todos[k] = FormatTodo(todo)
				}
			}
		} else {
			todos[k] = FormatTodo(todo)
		}
	}

	return todos, nil
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
