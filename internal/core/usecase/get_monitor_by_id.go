package usecase

import (
	"context"
	"strconv"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/exceptions"
)

type GetMonitorByIdUseCase struct {
	monitorRepository      domain.MonitorRepository
	notificationRepository domain.NotificationRepository
}

func NewGetMonitorByIdUseCase(
	monitorRepository domain.MonitorRepository,
	notificationRepository domain.NotificationRepository,
) *GetMonitorByIdUseCase {
	return &GetMonitorByIdUseCase{
		monitorRepository:      monitorRepository,
		notificationRepository: notificationRepository,
	}
}

func (g *GetMonitorByIdUseCase) Execute(context context.Context, id int64) (*domain.Monitor, error) {
	m, err := g.monitorRepository.GetById(context, id)

	if err != nil {
		return nil, exceptions.EntityNotFound("monitor", strconv.FormatInt(id, 10))
	}

	notifications, err := g.notificationRepository.GetByMonitorId(context, id)

	if err != nil {
		return nil, err
	}

	m.Notifications = notifications

	return m, nil
}
