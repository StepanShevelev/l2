package service

import (
	"github.com/StepanShevelev/l2/develop/timetable"
	"github.com/StepanShevelev/l2/develop/timetable/pkg/repository"
)

type EventService struct {
	repo repository.Calendar
}

func NewCalendarService(repo repository.Calendar) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event *timetable.Event) (int, error) {
	return s.repo.CreateEvent(event)
}
func (s *EventService) UpdateEvent(event *timetable.Event) error {
	return s.repo.UpdateEvent(event)
}
func (s *EventService) DeleteEvent(eventId int) error {
	return s.repo.DeleteEvent(eventId)
}
func (s *EventService) EventsForDay(event *timetable.Event) ([]timetable.ResultEvent, error) {
	return s.repo.EventsForDay(event)
}
func (s *EventService) EventsForWeek(event *timetable.Event) ([]timetable.ResultEvent, error) {
	return s.repo.EventsForWeek(event)
}
func (s *EventService) EventsForMonth(event *timetable.Event) ([]timetable.ResultEvent, error) {
	return s.repo.EventsForMonth(event)
}
