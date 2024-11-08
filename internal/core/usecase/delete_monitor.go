package usecase

import (
	"context"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/exceptions"
)

type DeleteMonitorUseCase struct {
	getMonitorById    *GetMonitorByIdUseCase
	monitorRepository domain.MonitorRepository
}

func NewDeleteMonitorUseCase(
	getMonitorById *GetMonitorByIdUseCase,
	monitorRepository domain.MonitorRepository) *DeleteMonitorUseCase {
	return &DeleteMonitorUseCase{
		getMonitorById:    getMonitorById,
		monitorRepository: monitorRepository,
	}
}

func (d *DeleteMonitorUseCase) Execute(ctx context.Context, id int64) error {
	_, err := d.getMonitorById.Execute(ctx, id)

	if err != nil {
		return exceptions.EntityNotFound("monitor", id)
	}

	_, err = d.monitorRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
