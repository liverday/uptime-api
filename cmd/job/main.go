package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"uptime-api/m/v2/cmd/api/deps"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/usecase"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Periodicity not found, should use one of these: %+v", domain.Periods)
	}

	p := domain.Periodicity(os.Args[1])

	if !slices.Contains(domain.Periods, p) {
		log.Fatalf("Wrong %s periodicity, should use one of these: %+v", p, domain.Periods)
	}

	log.Printf("Creating job: %s\n", p)
	d := deps.NewDependencies()
	gp := usecase.NewGetMonitorByPeriodicityUseCase(d.MonitorRepository)
	cr := usecase.NewCreateMonitorRunUseCase(d.MonitorRunRepository)
	us := usecase.NewUpdateMonitorStatusUseCase(d.MonitorRepository)
	tr := usecase.NewTriggerNotificationUseCase(
		d.NotificationRepository,
		d.MonitorRepository,
		d.CacheProvider,
	)

	c := cron.New()
	err := c.AddJob(fmt.Sprintf("@every %s", p.String()), NewChecker(
		p,
		gp,
		cr,
		us,
		tr))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Job added successfully, running it now")
	c.Start()
	log.Printf("Job started sucessfully. Fetching monitors to run a ping through them every: %s\n", p)

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q

	fmt.Println("Stopping program")
}
