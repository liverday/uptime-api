package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/exceptions"
	"uptime-api/m/v2/internal/core/inputs"
)

type UpdateMonitorUseCase struct {
	getMonitorById    *GetMonitorByIdUseCase
	monitorRepository domain.MonitorRepository
}

func NewUpdateMonitorUseCase(
	getMonitorById *GetMonitorByIdUseCase,
	monitorRepository domain.MonitorRepository,
) *UpdateMonitorUseCase {
	return &UpdateMonitorUseCase{
		getMonitorById:    getMonitorById,
		monitorRepository: monitorRepository,
	}
}

func (u *UpdateMonitorUseCase) Execute(ctx context.Context, input *inputs.UpdateMonitorInput) (*domain.Monitor, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	monitor, err := u.getMonitorById.Execute(ctx, input.MonitorId)

	if err != nil {
		return nil, exceptions.EntityNotFound("monitor", input.MonitorId)
	}

	err = monitor.Update(input)

	if err != nil {
		return nil, err
	}

	return u.monitorRepository.Update(ctx, monitor)
}
