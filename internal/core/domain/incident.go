package domain

import "time"

type Incident struct {
	Id             string    `json:"id" db:"id"`
	MonitorId      string    `json:"monitor_id" db:"monitor_id"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	Status         string    `json:"status" db:"status"`
	StartedAt      time.Time `json:"started_at" db:"started_at"`
	ResolvedAt     time.Time `json:"resolved_at" db:"resolved_at"`
	ResolvedBy     string    `json:"resolved_by" db:"resolved_by"`
	AcknowledgedBy string    `json:"acknowledged_by" db:"acknowledged_by"`
	AcknowledgedAt time.Time `json:"acknowledged_at" db:"acknowledged_at"`
}
