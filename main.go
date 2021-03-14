package main

import (
	"github.com/3ZA/tap-tap-track/html"
	"net/http"
)

func main() {
	http.HandleFunc("/activity", activity)
	http.ListenAndServe(":8585", nil)
}

func activity(w http.ResponseWriter, r *http.Request) {
	param := html.ActivityParams{
		Title: "Activity",
		ActivityHistory: map[string][]string{
			"cycle": []string{"mon", "tue", "wed"},
			"run":   []string{"wed", "sun"},
		},
	}
	html.Activity(w, param)
}
