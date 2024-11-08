package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
)

type AssignMonitorToNotificationUseCase struct {
	getMonitorById         *GetMonitorByIdUseCase
	notificationRepository domain.NotificationRepository
}

func NewAssignMonitorToNotificationUseCase(getMonitorById *GetMonitorByIdUseCase, notificationRepository domain.NotificationRepository) *AssignMonitorToNotificationUseCase {
	return &AssignMonitorToNotificationUseCase{
		getMonitorById:         getMonitorById,
		notificationRepository: notificationRepository,
	}
}

func (a *AssignMonitorToNotificationUseCase) Execute(ctx context.Context, input *inputs.AssignMonitorToNotificationInput) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)

	if err != nil {
		return err
	}

	_, err = a.getMonitorById.Execute(ctx, input.MonitorId)

	if err != nil {
		return err
	}

	return a.notificationRepository.AssignToMonitor(ctx, input.NotificationId, input.MonitorId)
}
