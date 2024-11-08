package usecase

import (
	"context"
	"errors"
	"uptime-api/m/v2/internal/core/domain"
)

type GetAllMonitorsByIdsUseCase struct {
	monitorRepository domain.MonitorRepository
}

func NewGetAllMonitorsByIdsUseCase(monitorRepository domain.MonitorRepository) *GetAllMonitorsByIdsUseCase {
	return &GetAllMonitorsByIdsUseCase{
		monitorRepository: monitorRepository,
	}
}

func (g *GetAllMonitorsByIdsUseCase) Execute(ctx context.Context, ids []int64) ([]*domain.Monitor, error) {
	if len(ids) == 0 {
		return nil, errors.New("at least one id is required")
	}

	monitors, err := g.monitorRepository.GetAllByIds(ctx, ids)

	if err != nil {
		return nil, err
	}

	return monitors, nil
}
