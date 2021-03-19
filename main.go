package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/3ZA/tap-tap-track/html"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	habitHandler := &HabitHandler{
		db: NewInMemoryStore(),
	}
	http.Handle("/habits", habitHandler)
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

type Store interface {
	GetHabitsByDate(date string) []*html.Habit
	Update(date string, habit *html.Habit)
}

type HabitHandler struct {
	db Store
}

func (*HabitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		now := time.Now()
		date := fmt.Sprintf("%02d %s %d", now.Day(), now.Month().String()[:3], now.Year())
		param := html.HabitsParams{
			Title:  "Habits",
			Date:   date,
			Habits: []string{"cycle", "run"},
		}
		html.Habits(w, param)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		habit := r.FormValue("habit")
		done := r.Form.Get("done")
		fmt.Printf("%s %s", string(habit), string(done))
	}
}
