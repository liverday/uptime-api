package main

import (
	"context"
	"crypto/tls"
	"github.com/panjf2000/ants/v2"
	"log"
	"net/http"
	"sync"
	"time"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
	"uptime-api/m/v2/internal/core/usecase"
)

type Checker struct {
	periodicity             domain.Periodicity
	getMonitorByPeriodicity *usecase.GetMonitorByPeriodicityUseCase
	createMonitorRun        *usecase.CreateMonitorRunUseCase
	updateMonitorStatus     *usecase.UpdateMonitorStatusUseCase
	triggerNotification     *usecase.TriggerNotificationUseCase
}

func NewChecker(
	p domain.Periodicity,
	g *usecase.GetMonitorByPeriodicityUseCase,
	c *usecase.CreateMonitorRunUseCase,
	u *usecase.UpdateMonitorStatusUseCase,
	t *usecase.TriggerNotificationUseCase,
) *Checker {
	return &Checker{
		periodicity:             p,
		getMonitorByPeriodicity: g,
		createMonitorRun:        c,
		updateMonitorStatus:     u,
		triggerNotification:     t,
	}
}

const poolSize = 5

var ctx = context.Background()

func (c *Checker) Run() {
	log.Printf("Starting job, looking for monitors with %s periodicity\n", string(c.periodicity))
	ms, err := c.getMonitorByPeriodicity.Execute(ctx, c.periodicity)

	if err != nil {
		log.Printf("error getting monitors with %s periodicity: %w", string(c.periodicity), err)
		return
	}

	log.Printf("Found %d monitors to check\n", len(ms))

	pool, _ := ants.NewPool(poolSize)

	defer pool.Release()

	var wg sync.WaitGroup
	for _, m := range ms {
		_ = pool.Submit(c.check(m, &wg))
	}
	wg.Wait()
}

func (c *Checker) check(m *domain.Monitor, wg *sync.WaitGroup) func() {
	wg.Add(1)
	return func() {
		log.Printf("Performing ping to %s for monitor with id: %d\n", m.Url, m.Id)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{
			Timeout:   45 * time.Second,
			Transport: tr,
		}

		input := HttpCheckerRequest{
			MonitorID: m.Id,
			URL:       m.Url,
			Method:    m.Method,
			Body:      m.Body,
		}

		r, err := performHttp(ctx, client, input)

		if err != nil {
			log.Printf("error performing ping for Monitor %d: %w", m.Id, err)
		}

		var status domain.MonitorStatus
		isSuccessful := r.Status >= 200 && r.Status < 300

		switch {
		case isSuccessful && r.Latency > m.DegradedAfter:
			status = domain.MonitorDegraded
		case !isSuccessful:
			status = domain.MonitorAlerting
		default:
			status = domain.MonitorActive
		}

		if m.Status != status {
			// new status is different than previous
			err = c.updateMonitorStatus.Execute(ctx, m.Id, status)

			if err != nil {
				log.Printf("error updating monitor status for Monitor %d: %w", m.Id, err)
			}

			err = c.triggerNotification.Execute(inputs.TriggerNotificationInput{
				MonitorId:  m.Id,
				Status:     string(status),
				StatusCode: r.Status,
				Message:    r.Error,
				Timing:     r.Timing,
			})

			if err != nil {
				log.Printf("error while triggering notification for Monitor %d: %w", m.Id, err)
			}
		}

		mr := &domain.MonitorRun{
			MonitorId:             m.Id,
			Status:                status,
			Latency:               r.Latency,
			RanAt:                 r.Timestamp,
			DnsStartedAt:          time.UnixMilli(r.Timing.DnsStart),
			DnsEndedAt:            time.UnixMilli(r.Timing.DnsDone),
			ConnectStartedAt:      time.UnixMilli(r.Timing.ConnectStart),
			ConnectEndedAt:        time.UnixMilli(r.Timing.ConnectDone),
			TlsHandshakeStartedAt: time.UnixMilli(r.Timing.TlsHandshakeStart),
			TlsHandshakeEndedAt:   time.UnixMilli(r.Timing.TlsHandshakeDone),
		}

		_, err = c.createMonitorRun.Execute(ctx, mr)

		if err != nil {
			log.Printf("error performing ping for Monitor %d: %w", m.Id, err)
		} else {
			log.Printf("Ping to %s for monitor with id: %d finished with status code: %d\n", m.Url, m.Id, r.Status)
			log.Printf("Timing details: %+v", r.Timing)
		}

		defer wg.Done()
	}
}
