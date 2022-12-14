package repository

import (
	"fmt"

	"github.com/StepanShevelev/l2/develop/timetable"
	"github.com/jmoiron/sqlx"
)

const eventsTable = "events"

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (r *EventPostgres) CreateEvent(event *timetable.Event) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createEventQuery := fmt.Sprintf("INSERT INTO %s (title, description, date) VALUES ($1, $2, $3) RETURNING id", eventsTable)
	row := tx.QueryRow(createEventQuery, event.Title, event.Description, event.ParsedDate)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *EventPostgres) UpdateEvent(event *timetable.Event) error {
	query := fmt.Sprintf("UPDATE %s SET title = $1, description = $2 WHERE id = $3", eventsTable)
	_, err := r.db.Exec(query, event.Title, event.Description, event.Id)

	return err
}

func (r *EventPostgres) DeleteEvent(eventId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", eventsTable)
	_, err := r.db.Exec(query, eventId)

	return err
}

func (r *EventPostgres) EventsForDay(event *timetable.Event) ([]timetable.ResultEvent, error) {
	var events []timetable.ResultEvent
	query := fmt.Sprintf("SELECT * FROM %s WHERE date BETWEEN $1 AND $2", eventsTable)

	err := r.db.Select(&events, query, event.MinDate, event.MaxDate)

	return events, err
}

func (r *EventPostgres) EventsForWeek(event *timetable.Event) ([]timetable.ResultEvent, error) {
	var events []timetable.ResultEvent
	query := fmt.Sprintf("SELECT * FROM %s WHERE date BETWEEN $1 AND $2", eventsTable)

	err := r.db.Select(&events, query, event.MinDate, event.MaxDate)

	return events, err
}

func (r *EventPostgres) EventsForMonth(event *timetable.Event) ([]timetable.ResultEvent, error) {
	var events []timetable.ResultEvent
	query := fmt.Sprintf("SELECT * FROM %s WHERE date BETWEEN $1 AND $2", eventsTable)

	err := r.db.Select(&events, query, event.MinDate, event.MaxDate)

	return events, err
}
