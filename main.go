package main

import (
	"github.com/3ZA/tap-tap-track/html"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/activity", activity)
	http.ListenAndServe(":8585", nil)
}

const sparkles = "✨"
const miss = "✗"

func activity(w http.ResponseWriter, r *http.Request) {
	param := html.ActivityParams{
		Title: "Activity",
		ActivityHistory: map[string][]string{
			"cycle": []string{sparkles, sparkles, miss, miss, miss, sparkles, sparkles},

			"run": []string{"wed", "sun"},
		},
	}
	html.Activity(w, param)
}
