package html

import (
	"embed"
	"io"
	"text/template"
)

//go:embed *
var files embed.FS

var (
	activity = parse("activity.html")
	habits   = parse("track.html")
)

type ActivityParams struct {
	Title           string
	ActivityHistory map[string][]string
}

type HabitsParams struct {
	Title  string
	Date   string
	Habits []string
}

func Habits(w io.Writer, p HabitsParams) error {
	return habits.Execute(w, p)
}

func Activity(w io.Writer, p ActivityParams) error {
	return activity.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}
