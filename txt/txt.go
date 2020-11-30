package txt

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Task struct {
	Key   int
	Value string
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		// panic(e)
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
		log.Fatal(err)
	}
	defer todoFile.Close()

	if _, err := todoFile.WriteString("\n" + task); err != nil {
		log.Fatal(err)
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

func ListTasks(queries []string) ([]string, error) {
	todoLines, err := linesInFile("todo.txt")
	if err != nil {
		return nil, err
	}

	var tasks []string

	for _, t := range todoLines {
		if len(queries) > 0 {
			for _, q := range queries {
				taskExists := contains(tasks, t)
				taskMatches := strings.Contains(t, q)
				if taskMatches && !taskExists {
					tasks = append(tasks, t)
				}
			}
		} else {
			tasks = append(tasks, t)
		}
	}

	return tasks, nil
}

func CreateTask(task string) error {
	if err := appendTodo(task); err != nil {
		log.Fatal(err)
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
	todoLines, err := linesInFile("todo.txt")
	if err != nil {
		return err
	}

	taskMap := make(map[int]string)
	for i, line := range todoLines {
		taskMap[i] = line
	}

	var completedTasks []string
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

func DeleteTask(key int) error {
	f, err := os.Open("todo.txt")
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var tasks []string

	for s.Scan() {
		tasks = append(tasks, s.Text())
	}

	tasks = append(tasks[:key-1], tasks[key:]...)

	if err := writeTodos(tasks); err != nil {
		return err
	}

	return nil
}
