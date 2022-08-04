package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/StepanShevelev/l2/develop/timetable"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

func getEventId(req *http.Request) (int, error) {
	var id *Id
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return 0, err
	}
	if err = json.Unmarshal(body, &id); err != nil {
		return 0, err
	}

	return id.Id, nil
}

func jsonParser(req *http.Request) (*timetable.Event, error) {
	var input *timetable.Event

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &input); err != nil {
		return nil, err
	}
	layout := "2006-01-02 15:04:05"
	input.ParsedDate, err = time.Parse(layout, input.Date)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func queryParser(req *http.Request) (*timetable.Event, error) {
	input := new(timetable.Event)
	u, err := url.Parse(req.RequestURI)
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	input.Date = q.Get("date")
	layout := "2006-01-02 15:04:05"
	input.ParsedDate, err = time.Parse(layout, input.Date)
	return input, nil
}

func eventsForDay(event *timetable.Event) {
	year, month, day := event.ParsedDate.Date()
	event.MinDate = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	event.MaxDate = time.Date(year, month, day, 23, 59, 59, 999999999, time.UTC)
}

func eventsForWeek(event *timetable.Event) {
	startOfWeek := WeekStart(event.ParsedDate.ISOWeek())
	event.MinDate = startOfWeek
	event.MaxDate = startOfWeek.AddDate(0, 0, 6)
}

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func eventsForMonth(event *timetable.Event) {
	year, month, _ := event.ParsedDate.Date()
	event.MinDate = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	event.MaxDate = time.Date(year, month+1, 0, 23, 59, 59, 999999999, time.UTC)
}
