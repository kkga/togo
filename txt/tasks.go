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
		fmt.Println(i+1, "|", task)
	}
	fmt.Println("---")
	fmt.Println("Total tasks: ", len(tasks))
}

func CreateTask(task string) error {
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	id, e := f.WriteString("\n" + task)
	fmt.Println(id)
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

// func itob(v int) []byte {
// 	b := make([]byte, 8)
// 	binary.BigEndian.PutUint64(b, uint64(v))
// 	return b
// }

// func btoi(b []byte) int {
// 	return int(binary.BigEndian.Uint64(b))
// }
