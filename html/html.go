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
)

type ActivityParams struct {
	Title           string
	ActivityHistory map[string][]string
}

func Activity(w io.Writer, p ActivityParams) error {
	return activity.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}
