package togo

import (
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
			"2020-01-30 todo with creation date and due date due:2020-02-02",
			Todo{
				Done:         false,
				CreationDate: time.Date(2020, 01, 30, 0, 0, 0, 0, time.UTC),
				DueDate:      time.Date(2020, 02, 02, 0, 0, 0, 0, time.UTC),
				Subject:      "todo with creation date and due date due:2020-02-02",
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
