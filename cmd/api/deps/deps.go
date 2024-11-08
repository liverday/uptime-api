package deps

import (
	"uptime-api/m/v2/internal/adapters/cache"
	"uptime-api/m/v2/internal/adapters/repository"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/ports"
)

type Dependencies struct {
	// MonitorRepository is the repository for the Monitor entity
	MonitorRepository      domain.MonitorRepository
	MonitorRunRepository   domain.MonitorRunRepository
	IncidentRepository     domain.IncidentRepository
	NotificationRepository domain.NotificationRepository
	StatusPageRepository   domain.StatusPageRepository
	CacheProvider          ports.CacheProvider
}

func NewDependencies() *Dependencies {
	db := openDatabase()

	return &Dependencies{
		MonitorRepository:      repository.NewMonitorRepository(db),
		MonitorRunRepository:   repository.NewMonitorRunRepository(db),
		NotificationRepository: repository.NewNotificationRepository(db),
		StatusPageRepository:   repository.NewStatusPageRepository(db),
		CacheProvider:          cache.NewCacheProvider(),
	}
}
