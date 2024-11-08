package domain

import (
	"errors"
	"slices"
	"time"
	"uptime-api/m/v2/internal/core/inputs"
)

type Monitor struct {
	Id            int64           `json:"id" db:"id"`
	Title         string          `json:"title" db:"title"`
	Description   string          `json:"description" db:"description"`
	Status        MonitorStatus   `json:"status" db:"status"`
	Periodicity   Periodicity     `json:"periodicity" db:"periodicity"`
	Url           string          `json:"url" db:"url"`
	Method        string          `json:"method" db:"method"`
	Headers       string          `json:"headers" db:"headers"`
	Body          string          `json:"body" db:"body"`
	DegradedAfter int64           `json:"degraded_after" db:"degraded_after"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" db:"updated_at"`
	Notifications []*Notification `json:"notifications" db:"-"`
}

type MonitorStatus string

const (
	MonitorActive   MonitorStatus = "active"
	MonitorAlerting MonitorStatus = "alerting"
	MonitorDegraded MonitorStatus = "degraded"
)

var AllStatus = []MonitorStatus{
	MonitorActive,
	MonitorAlerting,
	MonitorDegraded,
}

func NewMonitor(input *inputs.CreateMonitorInput) (*Monitor, error) {
	p := Periodicity(input.Periodicity)

	if !slices.Contains(Periods, p) {
		return nil, errors.New("invalid_periodicity")
	}

	m := &Monitor{
		Title:         input.Title,
		Description:   input.Description,
		Status:        MonitorActive,
		Periodicity:   p,
		Url:           input.Url,
		Method:        input.Method,
		Headers:       input.Headers,
		Body:          input.Body,
		DegradedAfter: input.DegradedAfter,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return m, nil
}

func (m *Monitor) Update(input *inputs.UpdateMonitorInput) error {
	p := Periodicity(input.Periodicity)

	if !slices.Contains(Periods, p) {
		return errors.New("invalid_periodicity")
	}

	m.Title = input.Title
	m.Description = input.Description
	m.Periodicity = p
	m.Url = input.Url
	m.Method = input.Method
	m.Headers = input.Headers
	m.Body = input.Body
	m.DegradedAfter = input.DegradedAfter

	return nil
}
