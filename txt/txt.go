package txt

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

func getTaskMap() (map[int]string, error) {
	m := make(map[int]string)
	todoLines, err := linesInFile("todo.txt")
	if err != nil {
		return nil, err
	}

	for i, line := range todoLines {
		m[i+1] = line
	}

	return m, nil

}

// GetTodoMap converts a todo.txt file into a map of Todos
func GetTodoMap() (map[int]Todo, error) {
	m := make(map[int]Todo)
	todoLines, err := linesInFile("todo.txt")
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

func GetTotalTodoLen(fileName string) (int, error) {
	lines, err := linesInFile(fileName)
	if err != nil {
		return 0, err
	}
	return len(lines), nil
}

func writeTodos(tasks []string) error {
	output := strings.Join(tasks, "\n")
	err := ioutil.WriteFile("todo.txt", []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
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

func ListTasks(queries []string) (map[int]string, error) {
	m, err := getTaskMap()
	if err != nil {
		return nil, err
	}

	tasks := make(map[int]string)

	for k, v := range m {
		if len(queries) > 0 {
			for _, q := range queries {
				_, exists := tasks[k]
				matches := strings.Contains(v, q)
				if !exists && matches {
					tasks[k] = v
				}
			}
		} else {
			tasks[k] = v
		}
	}

	return tasks, nil
}

func CreateTask(task string) error {
	if err := appendTodo(task); err != nil {
		return err
	}
	return nil
}

func CompleteTask(key int) (string, error) {
	todoLines, err := linesInFile("todo.txt")
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(todoLines[key-1], "x ") {
		todoLines[key-1] = strings.Replace(todoLines[key-1], "x ", "", 1)
	} else {
		todoLines[key-1] = fmt.Sprintf("x %s", todoLines[key-1])
	}

	completedTask := todoLines[key-1]

	if err := writeTodos(todoLines); err != nil {
		return "", err
	}

	return completedTask, nil
}

func ArchiveTasks() error {
	taskMap, err := getTaskMap()
	if err != nil {
		return err
	}

	completedTasks := make([]string, 0)
	for i, task := range taskMap {
		if strings.HasPrefix(task, "x ") {
			completedTasks = append(completedTasks, task)
			taskMap[i] = ""
		}
	}

	tasks := []string{}
	for _, task := range taskMap {
		if task != "" {
			tasks = append(tasks, task)
		}
	}

	if err := writeTodos(tasks); err != nil {
		return err
	}
	if err := appendDone(completedTasks); err != nil {
		return err
	}

	return nil
}

func DeleteTask(key int) (string, error) {
	taskMap, err := getTaskMap()
	if err != nil {
		return "", err
	}

	_, ok := taskMap[key]
	if !ok {
		return "", fmt.Errorf("Task #%d doesn't exist", key)
	}
	deletedTask := taskMap[key]
	delete(taskMap, key)

	tasks := []string{}
	for _, task := range taskMap {
		if task != "" {
			tasks = append(tasks, task)
		}
	}

	if err := writeTodos(tasks); err != nil {
		return "", err
	}

	return deletedTask, nil
}
