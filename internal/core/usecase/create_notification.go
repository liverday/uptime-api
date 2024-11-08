package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
)

type CreateNotificationUseCase struct {
	notificationRepository      domain.NotificationRepository
	assignMonitorToNotification *AssignMonitorToNotificationUseCase
}

func NewCreateNotificationUseCase(
	notificationRepository domain.NotificationRepository,
	assignMonitorToNotification *AssignMonitorToNotificationUseCase) *CreateNotificationUseCase {
	return &CreateNotificationUseCase{
		notificationRepository:      notificationRepository,
		assignMonitorToNotification: assignMonitorToNotification,
	}
}

func (c *CreateNotificationUseCase) Execute(ctx context.Context, input *inputs.CreateNotificationInput) (*domain.Notification, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	n := domain.NewNotification(input)

	n, err = c.notificationRepository.Save(ctx, n)

	if err != nil {
		return nil, err
	}

	for _, monitorId := range input.MonitorIds {
		err = c.assignMonitorToNotification.Execute(ctx, &inputs.AssignMonitorToNotificationInput{
			MonitorId:      monitorId,
			NotificationId: n.Id,
		})

		if err != nil {
			return nil, err
		}
	}

	return n, nil
}
