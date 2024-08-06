package out

import "queuecast/pkg/core/models"

type EventRepository interface {
	SaveEvent(event *models.Event) error
}
