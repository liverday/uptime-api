package usecase

import (
	"context"
	"fmt"
	"log"
	"time"
	"uptime-api/m/v2/internal/adapters/provider"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
	"uptime-api/m/v2/internal/core/ports"
)

type TriggerNotificationUseCase struct {
	notificationRepository domain.NotificationRepository
	monitorRepository      domain.MonitorRepository
	cacheProvider          ports.CacheProvider
}

func NewTriggerNotificationUseCase(
	notificationRepository domain.NotificationRepository,
	monitorRepository domain.MonitorRepository,
	cacheProvider ports.CacheProvider,
) *TriggerNotificationUseCase {
	return &TriggerNotificationUseCase{
		notificationRepository: notificationRepository,
		monitorRepository:      monitorRepository,
		cacheProvider:          cacheProvider,
	}
}

var ctx = context.Background()

func (t *TriggerNotificationUseCase) Execute(input inputs.TriggerNotificationInput) error {
	m, err := t.monitorRepository.GetById(ctx, input.MonitorId)
	n, err := t.notificationRepository.GetByMonitorId(ctx, input.MonitorId)

	if err != nil {
		return err
	}

	log.Printf("Found %d notifications to send\n", len(n))

	for _, notification := range n {
		log.Printf("Sending notification with provider: %s\n for monitor %d", notification.Provider, m.Id)
		p, err := provider.NewNotificationProvider(notification.Provider, notification.ProviderData)

		if err != nil {
			log.Printf("error creating notification provider: %s", notification.Provider)
			return err
		}

		for _, s := range domain.AllStatus {
			if string(s) == input.Status {
				continue
			}

			k := fmt.Sprintf("key:%d:%s:%s", m.Id, string(s), notification.Provider)
			err = t.cacheProvider.Delete(ctx, k)

			if err != nil {
				log.Printf("error deleting key: %s", k)
				return err
			}
		}

		k := fmt.Sprintf("key:%d:%s:%s", m.Id, input.Status, notification.Provider)

		result := t.cacheProvider.SetNX(ctx, k, "1", time.Minute*30)

		if !result {
			log.Printf("Notification already sent for monitor with id: %d and status: %s\n", m.Id, input.Status)
			continue
		}

		if input.Status == string(domain.MonitorDegraded) {
			err = p.SendDegraded(m, input.StatusCode, input.Timing)
		} else if input.Status == string(domain.MonitorActive) {
			err = p.SendRecovered(m, input.StatusCode, input.Timing)
		} else if input.Status == string(domain.MonitorAlerting) {
			err = p.SendAlerting(m, input.StatusCode, input.Message)
		}

		if err != nil {
			return err
		}

	}

	return nil
}
