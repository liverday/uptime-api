package usecase

import (
	"context"
	"uptime-api/m/v2/internal/core/domain"
)

type GetStatusPageByIdUseCase struct {
	StatusPageRepository domain.StatusPageRepository
	MonitorRepository    domain.MonitorRepository
}

func NewGetStatusPageByIdUseCase(
	statusPageRepository domain.StatusPageRepository,
	monitorRepository domain.MonitorRepository,
) *GetStatusPageByIdUseCase {
	return &GetStatusPageByIdUseCase{
		StatusPageRepository: statusPageRepository,
		MonitorRepository:    monitorRepository,
	}
}

func (g *GetStatusPageByIdUseCase) Execute(ctx context.Context, id int64) (*domain.StatusPage, error) {
	sp, err := g.StatusPageRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	ms, err := g.MonitorRepository.GetAllByStatusPageId(ctx, sp.Id)

	if err != nil {
		return nil, err
	}

	sp.Monitors = ms

	return sp, nil
}
