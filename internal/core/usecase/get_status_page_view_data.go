package usecase

import (
	"context"
	"math"
	"uptime-api/m/v2/internal/core/domain"
)

type GetStatusPageViewDataOutput struct {
	Title       string
	Description string
	Monitors    []*MonitorUptime
}

type MonitorUptime struct {
	Monitor *domain.Monitor
	Uptime  float64
	Entries []*domain.UptimeEntry
}

type GetStatusPageViewDataUseCase struct {
	getStatusPageById        *GetStatusPageByIdUseCase
	getMonitorRunsTimeSeries *GetMonitorRunsTimeSeriesUseCase
}

func NewGetStatusPageViewDataUseCase(
	getStatusPageById *GetStatusPageByIdUseCase,
	getMonitorRunsTimeSeries *GetMonitorRunsTimeSeriesUseCase,
) *GetStatusPageViewDataUseCase {
	return &GetStatusPageViewDataUseCase{
		getStatusPageById:        getStatusPageById,
		getMonitorRunsTimeSeries: getMonitorRunsTimeSeries,
	}
}

func calculateUptime(entries []*domain.UptimeEntry) float64 {
	var total int64
	var ok int64

	for _, e := range entries {
		total += e.Total
		ok += e.Ok
	}

	return math.Round((float64(ok)/float64(total))*10000) / 100
}

func (g *GetStatusPageViewDataUseCase) Execute(ctx context.Context, statusPageId int64) (*GetStatusPageViewDataOutput, error) {
	sp, err := g.getStatusPageById.Execute(ctx, statusPageId)

	if err != nil {
		return nil, err
	}

	var monitorUptime []*MonitorUptime

	for _, m := range sp.Monitors {
		entries, err := g.getMonitorRunsTimeSeries.Execute(ctx, m.Id)

		if err != nil {
			return nil, err
		}

		uptime := calculateUptime(entries)

		monitorUptime = append(monitorUptime, &MonitorUptime{
			Uptime:  uptime,
			Monitor: m,
			Entries: entries})
	}

	output := &GetStatusPageViewDataOutput{
		Title:       sp.Title,
		Description: sp.Description,
		Monitors:    monitorUptime,
	}

	return output, nil
}
