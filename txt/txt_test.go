package txt

import (
	"testing"

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
