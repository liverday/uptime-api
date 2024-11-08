package domain

import (
	"context"
)

type MonitorRunRepository interface {
	Save(ctx context.Context, mr *MonitorRun) (*MonitorRun, error)
	GetById(ctx context.Context, id int64) (*MonitorRun, error)
	GetMonitorRunsTimeSeries90d(ctx context.Context, monitorId int64) ([]*UptimeEntry, error)
}
