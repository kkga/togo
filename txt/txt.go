package txt

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Todo struct {
	Num  int
	Done bool
	Subj string
	Proj string
}

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
