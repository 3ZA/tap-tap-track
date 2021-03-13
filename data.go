package main

import "time"

type Activity struct {
	Name           string
	DatesCompleted []JSONDate
}

type Tracker struct {
	FName           string
	ActivityHistory map[string]*Activity
}

type JSONDate struct {
	time.Time
}

func (t JSONDate) MarshalJSON() ([]byte, error) {
	// YYYY-MM-DD
	return []byte(t.Format("2005-01-02")), nil
}

// Track stores todays date in the history of an activity.
// The presence of a date in the history indicates that the task was done on that date.
func (t *Tracker) Track(activityName string) {
	activity, ok := t.ActivityHistory[activityName]
	if !ok {
		activity = &Activity{Name: activityName}
		t.ActivityHistory[activityName] = activity
	}
	activity.track()
}

func (a *Activity) track() {
	a.DatesCompleted = append(a.DatesCompleted, JSONDate{time.Now()})
}
