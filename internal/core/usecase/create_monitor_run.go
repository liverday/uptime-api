package usecase

import (
	"context"
	"uptime-api/m/v2/internal/core/domain"
)

type CreateMonitorRunUseCase struct {
	monitorRunRepository domain.MonitorRunRepository
}

func NewCreateMonitorRunUseCase(m domain.MonitorRunRepository) *CreateMonitorRunUseCase {
	return &CreateMonitorRunUseCase{monitorRunRepository: m}
}

func (c *CreateMonitorRunUseCase) Execute(ctx context.Context, mr *domain.MonitorRun) (*domain.MonitorRun, error) {
	mr, err := c.monitorRunRepository.Save(ctx, mr)

	if err != nil {
		return nil, err
	}

	return mr, nil
}
