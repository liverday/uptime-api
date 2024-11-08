package domain

import (
	"time"
	"uptime-api/m/v2/internal/core/inputs"
)

type Notification struct {
	Id           int64                  `json:"id" db:"id"`
	Title        string                 `json:"title" db:"title"`
	Description  string                 `json:"description" db:"description"`
	Provider     string                 `json:"provider" db:"provider"`
	ProviderData map[string]interface{} `json:"provider_data" db:"provider_data"`
	CreatedAt    time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at" db:"updated_at"`
}

func NewNotification(input *inputs.CreateNotificationInput) *Notification {
	return &Notification{
		Title:        input.Title,
		Description:  input.Description,
		Provider:     input.Provider,
		ProviderData: input.ProviderData,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
