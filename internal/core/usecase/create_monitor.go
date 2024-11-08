package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
)

type CreateMonitorUseCase struct {
	monitorRepository domain.MonitorRepository
}

func NewCreateMonitorUseCase(monitorRepository domain.MonitorRepository) *CreateMonitorUseCase {
	return &CreateMonitorUseCase{monitorRepository: monitorRepository}
}

func (c *CreateMonitorUseCase) Execute(ctx context.Context, input *inputs.CreateMonitorInput) (*domain.Monitor, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	m, err := domain.NewMonitor(input)
	if err != nil {
		return nil, err
	}

	m, err = c.monitorRepository.Save(ctx, m)

	if err != nil {
		return nil, err
	}

	return m, nil
}
