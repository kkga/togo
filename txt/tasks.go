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

func AllTasks() ([]string, error) {
	f, err := os.Open("todo.txt")
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var tasks []string

	for s.Scan() {
		tasks = append(tasks, s.Text())
	}

	return tasks, nil
}

func CreateTask(task string) error {
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	_, e := f.WriteString(task + "\n")
	if e != nil {
		return e
	}
	return nil
}

func CompleteTask(key int) error {
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

	output := strings.Join(tasks, "\n")
	err = ioutil.WriteFile("todo.txt", []byte(output), 0644)
	check(err)

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
