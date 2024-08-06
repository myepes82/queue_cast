package in

import (
	"time"
)

type SaveEventCommandDto struct {
	Time    time.Time
	Origin  string
	Content string
}

type SaveEventCommand interface {
	Execute(event SaveEventCommandDto) error
}
