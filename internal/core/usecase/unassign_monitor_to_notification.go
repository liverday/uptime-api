package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
)

type UnassignMonitorFromNotificationUseCase struct {
	notificationRepository domain.NotificationRepository
}

func NewUnassignMonitorFromNotificationUseCase(notificationRepository domain.NotificationRepository) *UnassignMonitorFromNotificationUseCase {
	return &UnassignMonitorFromNotificationUseCase{
		notificationRepository: notificationRepository,
	}
}

func (u *UnassignMonitorFromNotificationUseCase) Execute(ctx context.Context, input *inputs.UnassignMonitorFromNotificationInput) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)

	if err != nil {
		return err
	}

	return u.notificationRepository.UnassignToMonitor(ctx, input.NotificationId, input.MonitorId)
}
