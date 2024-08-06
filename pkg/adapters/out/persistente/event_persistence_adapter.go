package persistence

import (
	"context"
	"fmt"
	"queuecast/pkg/core/models"
	"time"

	"go.uber.org/zap"
)

type EventPersistenceAdapter struct {
	db     *RedisPersistence
	logger *zap.Logger
}

func NewEventPersistenceAdapter(logger *zap.Logger, db *RedisPersistence) *EventPersistenceAdapter {
	return &EventPersistenceAdapter{
		logger: logger,
		db:     db,
	}
}

func (epa *EventPersistenceAdapter) SaveEvent(event *models.Event) error {

	eventData := map[string]interface{}{
		"content":    event.Content,
		"origin":     event.Origin,
		"created_at": event.Time.Format(time.RFC3339),
	}

	eventId := fmt.Sprintf("event:%s", event.Id)

	err := epa.db.session.HSet(context.Background(), eventId, eventData).Err()
	if err != nil {
		epa.logger.Error("failed to save event", zap.String("eventId", eventId), zap.Error(err))
		return err
	}

	epa.logger.Info("Event saved successfully", zap.String("event_id", eventId))
	return nil
}
