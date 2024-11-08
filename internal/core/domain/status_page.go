package domain

import (
	"time"
	"uptime-api/m/v2/internal/core/inputs"
)

type StatusPage struct {
	Id          int64      `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	Monitors    []*Monitor `json:"monitors" db:"-"`
}

func NewStatusPage(input *inputs.CreateStatusPageInput) *StatusPage {
	return &StatusPage{
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
