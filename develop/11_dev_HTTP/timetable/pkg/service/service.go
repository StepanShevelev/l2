package service

import (
	"github.com/StepanShevelev/l2/develop/timetable"
	"github.com/StepanShevelev/l2/develop/timetable/pkg/repository"
)

type Event interface {
	CreateEvent(event *timetable.Event) (int, error)
	UpdateEvent(event *timetable.Event) error
	DeleteEvent(eventId int) error
	EventsForDay(event *timetable.Event) ([]timetable.ResultEvent, error)
	EventsForWeek(event *timetable.Event) ([]timetable.ResultEvent, error)
	EventsForMonth(event *timetable.Event) ([]timetable.ResultEvent, error)
}

type Service struct {
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Event: NewCalendarService(repos.Timetable),
	}
}
