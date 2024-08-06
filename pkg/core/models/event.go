package models

import "time"

type Event struct {
	Id      string
	Time    time.Time
	Origin  string
	Content string
}

func NewEvent(Id string, Time time.Time, origin string, content string) *Event {
	return &Event{
		Id:      Id,
		Time:    Time,
		Origin:  origin,
		Content: content,
	}
}
