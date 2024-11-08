package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
	"uptime-api/m/v2/internal/core/util"
)

type CreateStatusPageUseCase struct {
	statusPageRepository domain.StatusPageRepository
	monitorRepository    domain.MonitorRepository
}

func NewCreateStatusPageUseCase(
	statusPageRepository domain.StatusPageRepository,
	monitorRepository domain.MonitorRepository) *CreateStatusPageUseCase {
	return &CreateStatusPageUseCase{
		statusPageRepository: statusPageRepository,
		monitorRepository:    monitorRepository,
	}
}

func (c *CreateStatusPageUseCase) Execute(ctx context.Context, input *inputs.CreateStatusPageInput) (*domain.StatusPage, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	sp := domain.NewStatusPage(input)

	sp, err = c.statusPageRepository.Save(ctx, sp)

	if err != nil {
		return nil, err
	}

	ms, err := c.monitorRepository.GetAllByIds(ctx, input.MonitorIds)

	if err != nil {
		return nil, err
	}

	monitorIds := util.MapSlice(ms, func(m *domain.Monitor) int64 { return m.Id })

	err = c.statusPageRepository.AssignMonitors(ctx, sp.Id, monitorIds)

	if err != nil {
		return nil, err
	}

	sp.Monitors = ms

	return sp, nil
}
