package usecase

import (
	"context"
	"uptime-api/m/v2/internal/core/domain"
)

type GetMonitorByPeriodicityUseCase struct {
	monitorRepository domain.MonitorRepository
}

func NewGetMonitorByPeriodicityUseCase(monitorRepository domain.MonitorRepository) *GetMonitorByPeriodicityUseCase {
	return &GetMonitorByPeriodicityUseCase{monitorRepository: monitorRepository}
}

func (g *GetMonitorByPeriodicityUseCase) Execute(ctx context.Context, p domain.Periodicity) ([]*domain.Monitor, error) {
	ms, err := g.monitorRepository.GetAllByPeriodicity(ctx, p)

	if err != nil {
		return nil, err
	}

	return ms, nil
}
