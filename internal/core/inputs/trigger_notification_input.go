package inputs

import "uptime-api/m/v2/internal/core/util"

type TriggerNotificationInput struct {
	MonitorId  int64
	Status     string
	StatusCode int
	Message    string
	Timing     util.Timing
}
