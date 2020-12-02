package txt

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

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

// Complete marks the todo as complete or incomplete
func (t *Todo) Complete() bool {
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

// linesInFile scans the given fileName and returns a slice of strings for each line
func linesInFile(fileName string) ([]string, error) {
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

// GetTodoMap converts a todo.txt file into a map of Todos
func GetTodoMap(fileName string) (map[int]Todo, error) {
	m := make(map[int]Todo)
	todoLines, err := linesInFile(fileName)
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
	dateLayout := "2006-01-02"
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
		s = append(s, todo.CompletionDate.Format("2006-01-02"))
	}
	if !todo.CreationDate.IsZero() {
		s = append(s, todo.CreationDate.Format("2006-01-02"))
	}
	if todo.Subject != "" {
		s = append(s, todo.Subject)
	}

	return strings.Join(s, " ")
}

// ListTodos returns a map of formatted todo strings that match given queries
func ListTodos(queries []string, fileName string) (map[int]string, error) {
	m, err := GetTodoMap(fileName)
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
		log.Fatalf("failed creating file: %s", err)
	}
	defer f.Close()

	datawriter := bufio.NewWriter(f)
	for _, todo := range todos {
		_, _ = datawriter.WriteString(todo + "\n")
	}
	if err := datawriter.Flush(); err != nil {
		log.Fatal(err)
	}

	return nil
}

// CompleteTodo marks given todo number as complete and writes the update to file
func CompleteTodo(key int) (Todo, error) {
	todoMap, err := GetTodoMap("todo.txt")
	if err != nil {
		return Todo{}, err
	}
	if todo, ok := todoMap[key]; ok {
		todo.Complete()
		todoMap[key] = todo
	} else {
		return Todo{}, errors.New("Non-existing todo number")
	}

	if err := WriteTodoMap(todoMap, "todo.txt"); err != nil {
		return Todo{}, err
	}

	return todoMap[key], nil
}

func GetTotalTodoLen(fileName string) (int, error) {
	lines, err := linesInFile(fileName)
	if err != nil {
		return 0, err
	}
	return len(lines), nil
}

func appendTodo(task string) error {
	todoFile, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer todoFile.Close()

	if _, err := todoFile.WriteString("\n" + task); err != nil {
		return err
	}

	return nil
}

func appendDone(tasks []string) error {
	doneFile, err := os.OpenFile("done.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	for _, task := range tasks {
		_, err := doneFile.WriteString("\n" + task)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateTask(task string) error {
	if err := appendTodo(task); err != nil {
		return err
	}
	return nil
}

// func ArchiveTasks() error {
// 	taskMap, err := GetTodoMap("todo.txt")
// 	if err != nil {
// 		return err
// 	}

// 	completedTasks := make([]string, 0)
// 	for i, task := range taskMap {
// 		if strings.HasPrefix(task, "x ") {
// 			completedTasks = append(completedTasks, task)
// 			taskMap[i] = ""
// 		}
// 	}

// 	tasks := []string{}
// 	for _, task := range taskMap {
// 		if task != "" {
// 			tasks = append(tasks, task)
// 		}
// 	}

// 	if err := writeTodos(tasks); err != nil {
// 		return err
// 	}
// 	if err := appendDone(completedTasks); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func DeleteTask(key int) (string, error) {
// 	taskMap, err := GetTodoMap("todo.txt")
// 	if err != nil {
// 		return "", err
// 	}

// 	_, ok := taskMap[key]
// 	if !ok {
// 		return "", fmt.Errorf("Task #%d doesn't exist", key)
// 	}
// 	deletedTask := taskMap[key]
// 	delete(taskMap, key)

// 	tasks := []string{}
// 	for _, task := range taskMap {
// 		if task.Subject != "" {
// 			tasks = append(tasks, task)
// 		}
// 	}

// 	if err := writeTodos(tasks); err != nil {
// 		return "", err
// 	}

// 	return deletedTask, nil
// }
