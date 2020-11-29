package txt

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

func linesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	result := []string{}

	for s.Scan() {
		result = append(result, s.Text())
	}

	return result
}

func ListTasks(queries []string) ([]string, error) {
	f, err := os.Open("todo.txt")
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var tasks []string

	for s.Scan() {
		if len(queries) > 0 {
			for _, q := range queries {
				taskExists := contains(tasks, s.Text())
				taskMatches := strings.Contains(s.Text(), q)
				if taskMatches && !taskExists {
					tasks = append(tasks, s.Text())
				}
			}
		} else {
			tasks = append(tasks, s.Text())
		}
	}

	return tasks, nil
}

func CreateTask(task string) error {
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	_, e := f.WriteString("\n" + task)
	if e != nil {
		return e
	}
	return nil
}

func CompleteTask(key int) (string, error) {
	f, err := os.Open("todo.txt")
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var tasks []string

	for s.Scan() {
		tasks = append(tasks, s.Text())
	}

	if strings.HasPrefix(tasks[key-1], "x ") {
		tasks[key-1] = strings.Replace(tasks[key-1], "x ", "", 1)
	} else {
		tasks[key-1] = fmt.Sprintf("x %s", tasks[key-1])
	}

	completedTask := tasks[key-1]

	output := strings.Join(tasks, "\n")
	err = ioutil.WriteFile("todo.txt", []byte(output), 0644)
	check(err)

	return completedTask, nil
}

func ArchiveTasks() error {
	lines := linesInFile("todo.txt")

	taskMap := make(map[int]string)
	var completedTasks []string

	for i, line := range lines {
		taskMap[i] = line
	}

	for i, task := range taskMap {
		if strings.HasPrefix(task, "x ") {
			completedTasks = append(completedTasks, task)
			taskMap[i] = ""
		}
		// if strings.HasPrefix(task, "x ") {
		// 	completedTasks = append(completedTasks, tasks[i])
		// 	tasks = append(tasks[:i], tasks[:i+1]...)
		// 	i = i - 1
		// }
	}
	fmt.Println(taskMap)

	output := []string{}

	for _, task := range taskMap {
		if task != "" {
			output = append(output, task)
		}
	}
	tasksOutput := strings.Join(output, "\n")
	err := ioutil.WriteFile("todo.txt", []byte(tasksOutput), 0644)
	if err != nil {
		return err
	}

	doneFile, err := os.OpenFile("done.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	for _, task := range completedTasks {
		_, err := doneFile.WriteString("\n" + task)
		if err != nil {
			return err
		}
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

	output := strings.Join(tasks, "\n")
	err = ioutil.WriteFile("todo.txt", []byte(output), 0644)
	check(err)

	return nil
}
