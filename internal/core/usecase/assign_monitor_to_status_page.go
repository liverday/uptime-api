package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/exceptions"
	"uptime-api/m/v2/internal/core/inputs"
	"uptime-api/m/v2/internal/core/util"
)

type AssignMonitorToStatusPageUseCase struct {
	getAllMonitorsByIds  *GetAllMonitorsByIdsUseCase
	getStatusPageById    *GetStatusPageByIdUseCase
	statusPageRepository domain.StatusPageRepository
}

func NewAssignMonitorToStatusPageUseCase(
	getAllMonitorsByIds *GetAllMonitorsByIdsUseCase,
	getStatusPageById *GetStatusPageByIdUseCase,
	statusPageRepository domain.StatusPageRepository) *AssignMonitorToStatusPageUseCase {
	return &AssignMonitorToStatusPageUseCase{
		getAllMonitorsByIds:  getAllMonitorsByIds,
		getStatusPageById:    getStatusPageById,
		statusPageRepository: statusPageRepository,
	}
}

func (a *AssignMonitorToStatusPageUseCase) Execute(ctx context.Context, input *inputs.AssignMonitorToStatusPageInput) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)

	if err != nil {
		return err
	}

	ms, err := a.getAllMonitorsByIds.Execute(ctx, input.MonitorIds)
	monitorIds := util.MapSlice(ms, func(m *domain.Monitor) int64 { return m.Id })

	if err != nil {
		return err
	}

	statusPage, err := a.getStatusPageById.Execute(ctx, input.StatusPageId)

	if err != nil {
		return exceptions.EntityNotFound("status page", input.StatusPageId)
	}

	err = a.statusPageRepository.AssignMonitors(ctx, statusPage.Id, monitorIds)

	if err != nil {
		return err
	}

	return nil
}
