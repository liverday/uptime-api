package inputs

type CreateMonitorInput struct {
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"`
	Periodicity   string `json:"periodicity" validate:"required"`
	Url           string `json:"url" validate:"required"`
	Method        string `json:"method"`
	Headers       string `json:"headers"`
	Body          string `json:"body"`
	DegradedAfter int64  `json:"degraded_after"`
}
