package definitions

import "queuecast/pkg/core/models"

type EventStore interface {
	SaveEvent(event models.Event) error
	GetEvents() ([]models.Event, error)
}
