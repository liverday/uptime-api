package usecase

import (
	"context"
	"strconv"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/exceptions"
)

type GetNotificationByIdUseCase struct {
	notificationRepository domain.NotificationRepository
}

func NewGetNotificationByIdUseCase(notificationRepository domain.NotificationRepository) *GetNotificationByIdUseCase {
	return &GetNotificationByIdUseCase{
		notificationRepository: notificationRepository,
	}
}

func (g *GetNotificationByIdUseCase) Execute(ctx context.Context, id int64) (*domain.Notification, error) {
	n, err := g.notificationRepository.GetById(ctx, id)

	if err != nil {
		return nil, exceptions.EntityNotFound("notification", strconv.FormatInt(id, 10))
	}

	return n, nil
}
