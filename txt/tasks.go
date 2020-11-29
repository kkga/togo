package txt

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Key   int
	Value string
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func AllTasks() {
	f, err := os.Open("todo.txt")
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var tasks []string

	for s.Scan() {
		tasks = append(tasks, s.Text())
	}

	for i, task := range tasks {
		number := ""
		if i+1 >= 10 {
			number = fmt.Sprintf("%d|", i+1)
		} else {
			number = fmt.Sprintf("%d |", i+1)
		}
		fmt.Println(fmt.Sprintf("%s %s", number, task))
	}
	fmt.Println("---")
	fmt.Println("Total tasks: ", len(tasks))
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

// func DeleteTask(key int) error {
// 	return db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket(taskBucket)
// 		return b.Delete(itob(key))
// 	})
// }
