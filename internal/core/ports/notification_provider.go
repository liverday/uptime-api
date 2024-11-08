package ports

import (
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/util"
)

type NotificationProvider interface {
	SendAlerting(m *domain.Monitor, statusCode int, message string) error
	SendDegraded(m *domain.Monitor, statusCode int, timing util.Timing) error
	SendRecovered(m *domain.Monitor, statusCode int, timing util.Timing) error
}
