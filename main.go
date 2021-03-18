package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/3ZA/tap-tap-track/html"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/activity", activity)
	http.HandleFunc("/habits", habits)
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

func habits(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	date := fmt.Sprintf("%02d %s %d", now.Day(), now.Month().String()[:3], now.Year())
	param := html.HabitsParams{
		Title:  "Habits",
		Date:   date,
		Habits: []string{"cycle", "run"},
	}
	html.Habits(w, param)
}
