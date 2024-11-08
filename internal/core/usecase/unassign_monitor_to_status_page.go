package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
)

type UnassignMonitorFromStatusPageUseCase struct {
	getMonitorById       *GetMonitorByIdUseCase
	getStatusPageById    *GetStatusPageByIdUseCase
	statusPageRepository domain.StatusPageRepository
}

func NewUnassignMonitorFromStatusPageUseCase(
	getMonitorById *GetMonitorByIdUseCase,
	getStatusPageById *GetStatusPageByIdUseCase,
	statusPageRepository domain.StatusPageRepository) *UnassignMonitorFromStatusPageUseCase {
	return &UnassignMonitorFromStatusPageUseCase{
		getMonitorById:       getMonitorById,
		getStatusPageById:    getStatusPageById,
		statusPageRepository: statusPageRepository,
	}
}

func (u *UnassignMonitorFromStatusPageUseCase) Execute(ctx context.Context, input *inputs.UnassignMonitorFromStatusPageInput) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)

	if err != nil {
		return err
	}

	return u.statusPageRepository.UnassignMonitors(ctx, input.StatusPageId, []int64{input.MonitorId})
}
