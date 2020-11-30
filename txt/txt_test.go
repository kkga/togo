package txt

import "testing"

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
		if !Equal(c.want, got) {
			t.Errorf("linesInFile(%q) == %q, WANT: %q", c.fileName, got, c.want)
		}
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
