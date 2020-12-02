package txt

import (
	"io/ioutil"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestLinesInFile(t *testing.T) {
	cases := []struct {
		fileName string
		want     []string
	}{
		{"todo-small.txt",
			[]string{
				"first todo",
				"last todo",
			}},
		{"todo.txt",
			[]string{
				"first todo",
				"2020-11-30 todo with date",
				"x 2020-11-30 2019-10-12 completed todo with date",
				"last todo",
			}},
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

func TestTodoComplete(t *testing.T) {
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

func TestFormatTodo(t *testing.T) {
	cases := []struct {
		todo Todo
		want string
	}{

		{Todo{
			Done:           true,
			CompletionDate: time.Date(2020, 05, 05, 0, 0, 0, 0, time.UTC),
			CreationDate:   time.Date(2020, 01, 12, 0, 0, 0, 0, time.UTC),
			Subject:        "completed todo with dates"},
			"x 2020-05-05 2020-01-12 completed todo with dates",
		},
		{Todo{
			Done:         false,
			CreationDate: time.Date(2019, 12, 01, 0, 0, 0, 0, time.UTC),
			Subject:      "todo with creation date"},
			"2019-12-01 todo with creation date",
		},
	}
	for _, c := range cases {
		got := FormatTodo(c.todo)
		if !cmp.Equal(c.want, got) {
			t.Errorf("WANT: %s, GOT: %s", c.want, got)
		}
	}
}

func TestListTodos(t *testing.T) {
	cases := []struct {
		fileName string
		want     map[int]string
	}{
		{"todo-small.txt",
			map[int]string{
				1: "first todo",
				2: "last todo",
			}},
		{"todo.txt",
			map[int]string{
				1: "first todo",
				2: "2020-11-30 todo with date",
				3: "x 2020-11-30 2019-10-12 completed todo with date",
				4: "last todo",
			}},
	}
	for _, c := range cases {
		got, err := ListTodos(make([]string, 0), "../testdata/"+c.fileName)
		if err != nil {
			t.Fatal(err)
		}
		if !cmp.Equal(c.want, got) {
			t.Errorf("%q == %q, WANT: %q", c.fileName, got, c.want)
		}
	}
}

func TestWriteTodoMap(t *testing.T) {
	cases := []struct {
		m             map[int]Todo
		goldenResPath string
	}{
		{
			map[int]Todo{
				1: {Done: false, Subject: "first todo"},
				2: {Done: true, Subject: "second todo"},
				3: {Done: false, Subject: "another todo"},
				4: {Done: true, Subject: "last todo"},
			},
			"write-test-1.golden",
		},
		{
			map[int]Todo{
				1: {
					Done:         false,
					Subject:      "first todo",
					CreationDate: time.Date(2020, 12, 12, 0, 0, 0, 0, time.UTC)},
				2: {
					Done:           true,
					Subject:        "completed todo with a +project",
					CreationDate:   time.Date(2019, 12, 12, 0, 0, 0, 0, time.UTC),
					CompletionDate: time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)},
				3: {
					Done:    false,
					Subject: "another todo"},
				4: {
					Done:    true,
					Subject: "last todo"},
			},
			"write-test-2.golden",
		},
	}
	for i, c := range cases {
		if err := WriteTodoMap(c.m, "../testdata/write-test-"+strconv.Itoa(i+1)+".output"); err != nil {
			t.Fatal(err)
		}
		want, err := ioutil.ReadFile("../testdata/write-test-" + strconv.Itoa(i+1) + ".golden")
		if err != nil {
			t.Fatal(err)
		}
		got, err := ioutil.ReadFile("../testdata/write-test-" + strconv.Itoa(i+1) + ".output")
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got) {
			t.Fatal("WriteTodoMap compare error")
		}
	}
}
