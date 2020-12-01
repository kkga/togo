package txt

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestLinesInFile(t *testing.T) {
	cases := []struct {
		fileName string
		want     []string
	}{
		{"todo-small.txt", []string{"first todo", "last todo"}},
		{"todo.txt", []string{"first todo", "2020-11-30 todo with date", "x 2020-11-30 completed todo with date", "last todo"}},
	}
	for _, c := range cases {
		got, err := linesInFile("../testdata/" + c.fileName)
		if err != nil {
			t.Fatal("Cannot read file")
		}
		if !cmp.Equal(c.want, got) {
			t.Errorf("linesInFile(%q) == %q, WANT: %q", c.fileName, got, c.want)
		}
	}
}

func TestTaskComplete(t *testing.T) {
	cases := []struct {
		todo Todo
		want Todo
	}{
		{
			Todo{Done: false, Subject: "my Todo"},
			Todo{Done: true, Subject: "my Todo"},
		},
		{
			Todo{Done: true, Subject: "my Todo"},
			Todo{Done: false, Subject: "my Todo"},
		},
	}
	for _, c := range cases {
		todo := c.todo
		todo.Complete()

		if !cmp.Equal(c.want, todo) {
			t.Errorf("GIVEN: %v, WANT: %v, GOT: %v", c.todo, c.want, todo)
		}
	}
}

func TestParseTodo(t *testing.T) {
	cases := []struct {
		todoStr string
		want    Todo
	}{
		{"some todo",
			Todo{
				Done:    false,
				Subject: "some todo"},
		},
		{"x some completed todo",
			Todo{
				Done:    true,
				Subject: "some completed todo"},
		},
		{"2020-01-30 todo with creation date",
			Todo{
				Done:         false,
				CreationDate: time.Date(2020, 01, 30, 0, 0, 0, 0, time.UTC),
				Subject:      "todo with creation date"},
		},
		{"x 2020-05-05 2020-01-12 completed todo with dates",
			Todo{
				Done:           true,
				CompletionDate: time.Date(2020, 05, 05, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 01, 12, 0, 0, 0, 0, time.UTC),
				Subject:        "completed todo with dates"}},
	}

	for _, c := range cases {
		got := ParseTodo(c.todoStr)
		if !cmp.Equal(got, c.want) {
			t.Errorf("WANT: %v, GOT: %v", c.want, got)
		}
	}
}
