package controllers

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type EventController struct {
	repo ports.EventRepository
}

func NewEventController(repo ports.EventRepository) *EventController {
	return &EventController{
		repo: repo,
	}
}

func (srv *EventController) Get(id uuid.UUID) (domain.Event, error) {
	event, err := srv.repo.Get(id)
	if err != nil {
		return domain.Event{}, errors.New("get failed")
	}
	return event, nil
}

func (srv *EventController) Create(locationId uuid.UUID, ControllerId uuid.UUID, resourceId uuid.UUID, customerId uuid.UUID, startAt time.Time, endAt time.Time, eventType string) (domain.Event, error) {
	event := domain.NewEvent(locationId, ControllerId, resourceId, customerId, startAt, endAt, eventType)

	if err := srv.repo.Save(event); err != nil {
		return domain.Event{}, errors.New("create event into repository has failed")
	}
	return event, nil
}
