package inputs

type CreateStatusPageInput struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description"`
	MonitorIds  []int64 `json:"monitor_ids" validate:"required"`
}
