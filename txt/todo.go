package txt

import (
	"regexp"
	"strings"
	"time"
)

// Todo represents a single todo item
type Todo struct {
	Done           bool
	Original       string // original raw text from todo.txt
	Priority       string
	CompletionDate time.Time
	CreationDate   time.Time
	DueDate        time.Time
	Subject        string
	Projects       []string
	Contexts       []string
	Tags           []string
}

// ToggleDone marks the todo as complete or incomplete
func (t *Todo) ToggleDone() bool {
	t.Done = !t.Done
	return t.Done
}

// ParseTodo converts a string into a Todo struct
func ParseTodo(todo string) Todo {
	var (
		done           bool
		original       string
		subject        string
		completionDate time.Time
		creationDate   time.Time
		dueDate        time.Time
		priority       string
		projects       []string
		contexts       []string
	)

	var (
		prioRe    = regexp.MustCompile(`\([A-Z]\)`)
		dateRe    = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
		dueDateRe = regexp.MustCompile(`due:\d{4}-\d{2}-\d{2}`)
		projectRe = regexp.MustCompile(`\+\w+`)
		contextRe = regexp.MustCompile(`\@\w+`)
	)

	original = todo

	// check if the task is done
	if strings.HasPrefix(todo, "x ") {
		done = true
		todo = strings.Replace(todo, "x ", "", 1)
	}

	// parse priority
	if priority = prioRe.FindString(todo); priority != "" {
		todo = strings.Replace(todo, priority+" ", "", 1)
		priority = strings.Replace(priority, "(", "", 1)
		priority = strings.Replace(priority, ")", "", 1)
	}

	// parse dates
	// TODO this currently parses all dates, but need to tweak the regex to only
	// parse dates at the beginning
	// maybe split into separate functions for parsing creationDate and completionDate
	dates := dateRe.FindAllString(todo, -1)

	if len(dates) == 1 {
		todo = strings.Replace(todo, dates[0]+" ", "", 1)
		date, _ := time.Parse(dateLayout, dates[0])
		creationDate = date
	} else if len(dates) > 1 {
		todo = strings.Replace(todo, dates[0]+" ", "", 1)
		compDate, _ := time.Parse(dateLayout, dates[0])
		completionDate = compDate
		todo = strings.Replace(todo, dates[1]+" ", "", 1)
		creatDate, _ := time.Parse(dateLayout, dates[1])
		creationDate = creatDate
	}

	// parse due date
	if d := dueDateRe.FindString(todo); d != "" {
		dueDate, _ = time.Parse(dateLayout, strings.Replace(d, "due:", "", 1))
	}

	// parse projects
	// TODO: make this work with +my-project
	if p := projectRe.FindAllString(todo, -1); len(p) > 0 {
		for _, v := range p {
			projects = append(projects, v)
		}
	}

	// parse contexts
	if c := contextRe.FindAllString(todo, -1); len(c) > 0 {
		for _, v := range c {
			contexts = append(contexts, v)
		}
	}

	subject = todo

	return Todo{
		Done:           done,
		Original:       original,
		Subject:        subject,
		CompletionDate: completionDate,
		CreationDate:   creationDate,
		DueDate:        dueDate,
		Priority:       priority,
		Projects:       projects,
		Contexts:       contexts,
	}
}

// FormatTodo converts a Todo struct into a todo.txt string for writing
// TODO turn this into a String() method on Todo
func FormatTodo(todo Todo) string {
	s := make([]string, 0)

	if todo.Done {
		s = append(s, "x")
	}
	if todo.Priority != "" {
		s = append(s, "("+todo.Priority+")")
	}
	if !todo.CompletionDate.IsZero() {
		s = append(s, todo.CompletionDate.Format(dateLayout))
	}
	if !todo.CreationDate.IsZero() {
		s = append(s, todo.CreationDate.Format(dateLayout))
	}
	if todo.Subject != "" {
		s = append(s, todo.Subject)
	}

	return strings.Join(s, " ")
}
