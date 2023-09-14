package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"event-service/pkg/customerror"
)

type Event struct {
	ID          string
	name        string
	startTime   *time.Time
	endTime     *time.Time
	location    string
	description string

	branch          *Branch
	attendant_count int
	attendants      *[]Attendant //optional
}

type EventDTO struct {
	ID          string
	Name        string
	StartTime   *time.Time
	EndTime     *time.Time
	Location    string
	Description string

	organizer *Branch
}

// mapping for DTO to Entity
func NewEvent(event EventDTO) (*Event, error) {
	if event.ID == "" || event.Name == "" || event.StartTime == nil || event.EndTime == nil || event.Location == "" || event.Description == "" {
		return nil, errors.New(customerror.ERROR_FIELD_ENTITY)
	}

	return &Event{
		ID:          event.ID,
		name:        event.Name,
		startTime:   event.StartTime,
		endTime:     event.EndTime,
		location:    event.Location,
		description: event.Description,
		organizer:   event.organizer,
	}, nil
}

// getter & setter for entity
func (e Event) SetID() {
	e.ID = uuid.New().String()
}

func (e *Event) GetID() string {
	return e.ID
}

func (e Event) SetName(name string) {
	e.name = name
}

func (e *Event) GetName() string {
	return e.name
}

func (e Event) SetStartTime(startTime *time.Time) {
	e.startTime = startTime
}

func (e *Event) GetStartTime() *time.Time {
	return e.startTime
}

func (e Event) SetEndTime(endTime *time.Time) {
	e.endTime = endTime
}

func (e *Event) GetEndTime() *time.Time {
	return e.endTime
}

func (e Event) SetLocation(location string) {
	e.location = location
}

func (e *Event) GetLocation() string {
	return e.location
}

func (e Event) SetDescription(description string) {
	e.description = description
}

func (e *Event) GetDescription() string {
	return e.description
}
