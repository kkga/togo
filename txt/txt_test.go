package txt

import (
	"io/ioutil"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestTodoToggleDone(t *testing.T) {
	tests := []struct {
		name string
		todo Todo
		want Todo
	}{
		{
			"toggle done",
			Todo{Done: false, Subject: "undone"},
			Todo{Done: true, Subject: "undone"},
		},
		{
			"toggle undone",
			Todo{Done: true, Subject: "done"},
			Todo{Done: false, Subject: "done"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.todo.ToggleDone()
			want := tt.want.Done

			if got != want {
				t.Errorf("GIVEN: %v, WANT: %v, GOT: %v", tt.todo, got, want)
			}
		})
	}
}

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
		got, err := LinesInFile("../testdata/" + c.fileName)
		if err != nil {
			t.Fatal("Cannot read file")
		}
		if !cmp.Equal(c.want, got) {
			t.Errorf("linesInFile(%q) == %q, WANT: %q", c.fileName, got, c.want)
		}
	}
}

func TestTodoMap(t *testing.T) {
	tests := []struct {
		fileName string
		want     map[int]Todo
	}{
		{
			"todo-small.txt",
			map[int]Todo{
				1: {Done: false, Subject: "first todo"},
				2: {Done: false, Subject: "last todo"},
			},
		},
		{
			"todo.txt",
			map[int]Todo{
				1: {Subject: "first todo"},
				2: {Subject: "todo with date", CreationDate: time.Date(2020, 11, 30, 0, 0, 0, 0, time.UTC)},
				3: {Done: true, Subject: "completed todo with date", CompletionDate: time.Date(2020, 11, 30, 0, 0, 0, 0, time.UTC), CreationDate: time.Date(2019, 10, 12, 0, 0, 0, 0, time.UTC)},
				4: {Subject: "last todo"},
			},
		},
	}
	for _, tt := range tests {
		name := tt.fileName
		t.Run(name, func(t *testing.T) {
			got, err := TodoMap("../testdata/" + tt.fileName)
			if err != nil {
				t.Fatal("Can't get TodoMap:", err)
			}
			if !cmp.Equal(tt.want, got) {
				t.Errorf("TodoMap mismatch (-want +got):\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestParseTodo(t *testing.T) {
	tests := []struct {
		todoStr string
		want    Todo
	}{
		{
			"some todo",
			Todo{
				Done:    false,
				Subject: "some todo",
			},
		},
		{
			"x some completed todo",
			Todo{
				Done:    true,
				Subject: "some completed todo",
			},
		},
		{
			"(A) todo with priority",
			Todo{
				Done:     false,
				Priority: "A",
				Subject:  "todo with priority",
			},
		},
		{
			"x (C) completed todo with priority",
			Todo{
				Done:     true,
				Priority: "C",
				Subject:  "completed todo with priority",
			},
		},
		{
			"2020-01-30 todo with creation date",
			Todo{
				Done:         false,
				CreationDate: time.Date(2020, 01, 30, 0, 0, 0, 0, time.UTC),
				Subject:      "todo with creation date",
			},
		},
		{
			"x 2020-05-05 2020-01-12 completed todo with dates",
			Todo{
				Done:           true,
				CompletionDate: time.Date(2020, 05, 05, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 01, 12, 0, 0, 0, 0, time.UTC),
				Subject:        "completed todo with dates",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.todoStr, func(t *testing.T) {
			got := ParseTodo(tt.todoStr)
			if !cmp.Equal(tt.want, got) {
				t.Errorf("%s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestFormatTodo(t *testing.T) {
	tests := []struct {
		todo Todo
		want string
	}{

		{
			Todo{
				Done:           true,
				CompletionDate: time.Date(2020, 05, 05, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 01, 12, 0, 0, 0, 0, time.UTC),
				Subject:        "completed todo with dates",
			},
			"x 2020-05-05 2020-01-12 completed todo with dates",
		},
		{
			Todo{
				Done:         false,
				CreationDate: time.Date(2019, 12, 01, 0, 0, 0, 0, time.UTC),
				Subject:      "todo with creation date",
			},
			"2019-12-01 todo with creation date",
		},
		{
			Todo{
				Done:         false,
				Priority:     "A",
				CreationDate: time.Date(2019, 12, 01, 0, 0, 0, 0, time.UTC),
				Subject:      "todo with prio and creation date",
			},
			"(A) 2019-12-01 todo with prio and creation date",
		},
		{
			Todo{
				Done:           true,
				Priority:       "C",
				CompletionDate: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2019, 12, 01, 0, 0, 0, 0, time.UTC),
				Subject:        "completed todo with prio and creation date",
			},
			"x (C) 2020-01-01 2019-12-01 completed todo with prio and creation date",
		},
	}
	for _, tt := range tests {
		name := tt.todo.Subject
		t.Run(name, func(t *testing.T) {
			got := FormatTodo(tt.todo)
			if !cmp.Equal(tt.want, got) {
				t.Errorf("%s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestWriteTodoMap(t *testing.T) {
	cases := []struct {
		m              map[int]Todo
		goldenFilePath string
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
