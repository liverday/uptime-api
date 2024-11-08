package inputs

type UpdateMonitorInput struct {
	MonitorId     int64
	Title         string `validate:"required"`
	Description   string
	Periodicity   string `validate:"required"`
	Url           string `validate:"required"`
	Method        string
	Headers       string
	Body          string
	DegradedAfter int64
}
