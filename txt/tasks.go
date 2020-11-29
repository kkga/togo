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

	for _, task := range tasks {
		fmt.Println(task)
	}
	fmt.Println("---")
	fmt.Println("Total tasks: ", len(tasks))
}

// func CreateTask(task string) (int, error) {
// 	var id int
// 	err := db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket(taskBucket)
// 		id64, _ := b.NextSequence()
// 		id = int(id64)
// 		key := itob(int(id64))
// 		return b.Put(key, []byte(task))
// 	})
// 	if err != nil {
// 		return -1, err
// 	}
// 	return id, nil
// }

// func AllTasks() ([]Task, error) {
// 	var tasks []Task
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket(taskBucket)
// 		c := b.Cursor()

// 		for k, v := c.First(); k != nil; k, v = c.Next() {
// 			tasks = append(tasks, Task{
// 				Key:   btoi(k),
// 				Value: string(v),
// 			})
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tasks, nil
// }

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
