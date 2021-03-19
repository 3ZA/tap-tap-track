package main

import "github.com/3ZA/tap-tap-track/html"

type InMemoryStore struct {
	// map<date, map<activity, done>>
	state map[string]map[string]bool
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		state: map[string]map[string]bool{},
	}
}

func (i *InMemoryStore) GetHabitsByDate(date string) []*html.Habit {
	result := []*html.Habit{}
	if habits, ok := i.state[date]; ok {
		for name, done := range habits {
			result = append(result, &html.Habit{Name: name, Done: done})
		}
	}
	return result
}

func (i *InMemoryStore) Update(date string, h *html.Habit) {
	if habits, ok := i.state[date]; ok {
		habits[h.Name] = h.Done
	} else {
		i.state[date] = map[string]bool{
			h.Name: h.Done,
		}
	}
}
