package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/3ZA/tap-tap-track/html"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	db, err := NewBoltDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	habitHandler := &HabitHandler{
		db: db,
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
	//GetHabitsByDate(date string) []*html.Habit
	//Update(date string, habit *html.Habit)
	RetrieveEntry(time.Time) (*Entry, error)
	AddEntry(*Entry) error
	viewAll()
}

type HabitHandler struct {
	db Store
}

func (h *HabitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		now := time.Now()
		entry, err := h.db.RetrieveEntry(now)
		if err != nil {
			log.Print(err)
			return
		}
		if entry == nil {
			param := html.HabitsParams{
				Title:  "Habits",
				Date:   todaysDate(),
				Habits: map[string]bool{"cycle": false, "run": false},
			}
			html.Habits(w, param)
		} else {
			habitMap := map[string]bool{}
			for name, record := range entry.Records {
				habitMap[name] = record.Done
			}
			param := html.HabitsParams{
				Title:  "Habits",
				Date:   todaysDate(),
				Habits: habitMap,
			}
			html.Habits(w, param)
		}

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		habit := r.FormValue("habit")
		done := r.Form.Get("done")
		habitDone, _ := strconv.ParseBool(done)

		now := time.Now()
		entry, err := h.db.RetrieveEntry(now)
		if err != nil {
			log.Print(err)
			return
		}
		if entry == nil {
			entry = &Entry{Date: now, Records: map[string]*ActivityRecord{}}
		}

		if record, ok := entry.Records[habit]; ok {
			record.Done = habitDone
		} else {
			entry.Records[habit] = &ActivityRecord{Name: habit, Done: habitDone}
		}
		err = h.db.AddEntry(entry)
		if err != nil {
			log.Print(err)
		}

		//h.db.viewAll()
	}
}

func todaysDate() string {
	now := time.Now()
	return fmt.Sprintf("%02d %s %d", now.Day(), now.Month().String()[:3], now.Year())
}

func timeToDateString(t time.Time) string {
	return fmt.Sprintf("%02d %s %d", t.Day(), t.Month().String()[:3], t.Year())
}
