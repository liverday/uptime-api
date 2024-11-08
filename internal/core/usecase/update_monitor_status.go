package usecase

import (
	"context"
	"uptime-api/m/v2/internal/core/domain"
)

type UpdateMonitorStatusUseCase struct {
	monitorRepository domain.MonitorRepository
}

func NewUpdateMonitorStatusUseCase(monitorRepository domain.MonitorRepository) *UpdateMonitorStatusUseCase {
	return &UpdateMonitorStatusUseCase{monitorRepository: monitorRepository}
}

func (u *UpdateMonitorStatusUseCase) Execute(ctx context.Context, monitorId int64, status domain.MonitorStatus) error {
	return u.monitorRepository.UpdateStatus(ctx, monitorId, status)
}
