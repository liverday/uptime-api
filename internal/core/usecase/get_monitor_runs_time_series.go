package usecase

import (
	"context"
	"uptime-api/m/v2/internal/core/domain"
)

type GetMonitorRunsTimeSeriesUseCase struct {
	monitorRunRepository domain.MonitorRunRepository
}

func NewGetMonitorRunsTimeSeriesUseCase(monitorRunRepository domain.MonitorRunRepository) *GetMonitorRunsTimeSeriesUseCase {
	return &GetMonitorRunsTimeSeriesUseCase{
		monitorRunRepository: monitorRunRepository,
	}
}

func (g *GetMonitorRunsTimeSeriesUseCase) Execute(ctx context.Context, monitorId int64) ([]*domain.UptimeEntry, error) {
	return g.monitorRunRepository.GetMonitorRunsTimeSeries90d(ctx, monitorId)
}
