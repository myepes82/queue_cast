package usecases

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"queuecast/pkg/core/models"
	"queuecast/pkg/port/in"
	"queuecast/pkg/port/out"
)

type SaveEventUseCase struct {
	eventRepository out.EventRepository
	logger          *zap.Logger
}

func NewSaveEventUseCase(eventRepository out.EventRepository, logger *zap.Logger) *SaveEventUseCase {
	return &SaveEventUseCase{
		eventRepository: eventRepository,
		logger:          logger,
	}
}

func (uc *SaveEventUseCase) Execute(eventDto in.SaveEventCommandDto) error {
	uc.logger.Info("saving event record")
	id, err := uuid.NewV7()
	if err != nil {
		uc.logger.Error("error generating uuid", zap.Error(err))
		return err
	}
	event := models.NewEvent(id.String(), eventDto.Time, eventDto.Origin, eventDto.Content)
	return uc.eventRepository.SaveEvent(event)
}
