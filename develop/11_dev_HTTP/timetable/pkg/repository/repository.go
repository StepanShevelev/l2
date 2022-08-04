package repository

import (
	"github.com/StepanShevelev/l2/develop/timetable"
	"github.com/jmoiron/sqlx"
)

type Timetable interface {
	CreateEvent(event *timetable.Event) (int, error)
	UpdateEvent(event *timetable.Event) error
	DeleteEvent(eventId int) error
	EventsForDay(event *timetable.Event) ([]timetable.ResultEvent, error)
	EventsForWeek(event *timetable.Event) ([]timetable.ResultEvent, error)
	EventsForMonth(event *timetable.Event) ([]timetable.ResultEvent, error)
}

type Repository struct {
	Timetable
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Timetable: NewEventPostgres(db),
	}
}
