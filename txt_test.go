package togo

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
		got, err := LinesInFile("testdata/" + c.fileName)
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
			got, err := TodoMap("testdata/" + tt.fileName)
			if err != nil {
				t.Fatal("Can't get TodoMap:", err)
			}
			if !cmp.Equal(tt.want, got) {
				t.Errorf("TodoMap mismatch (-want +got):\n%s", cmp.Diff(tt.want, got))
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
		if err := WriteTodoMap(c.m, "testdata/write-test-"+strconv.Itoa(i+1)+".output"); err != nil {
			t.Fatal(err)
		}
		want, err := ioutil.ReadFile("testdata/write-test-" + strconv.Itoa(i+1) + ".golden")
		if err != nil {
			t.Fatal(err)
		}
		got, err := ioutil.ReadFile("testdata/write-test-" + strconv.Itoa(i+1) + ".output")
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got) {
			t.Fatal("WriteTodoMap compare error")
		}
	}
}

func TestProjects(t *testing.T) {
	tests := []struct {
		m    map[int]Todo
		want []string
	}{
		{
			map[int]Todo{
				1: {Done: false, Subject: "first +project todo", Projects: []string{"project"}},
				2: {Done: false, Subject: "last todo +project", Projects: []string{"project"}},
			},
			[]string{"project"},
		},
		{
			map[int]Todo{
				1: {Done: false, Subject: "first +hey todo", Projects: []string{"hey"}},
				2: {Done: false, Subject: "last todo +ho", Projects: []string{"ho"}},
				3: {Done: false, Subject: "last todo +other +ho", Projects: []string{"other", "ho"}},
				4: {Done: false, Subject: "last todo +hey", Projects: []string{"hey"}},
				5: {Done: false, Subject: "last todo +wassup +hey", Projects: []string{"wassup", "hey"}},
			},
			[]string{"hey", "ho", "other", "wassup"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Projects(tt.m)
			if err != nil {
				t.Fatal("Can't get projects:", err)
			}
			if !cmp.Equal(tt.want, got) {
				t.Errorf("Projects mismatch:\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}
