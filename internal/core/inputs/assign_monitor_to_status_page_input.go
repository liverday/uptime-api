package inputs

type AssignMonitorToStatusPageInput struct {
	StatusPageId int64   `json:"status_page_id" validate:"required"`
	MonitorIds   []int64 `json:"monitor_id" validate:"min=0"`
}

type UnassignMonitorFromStatusPageInput struct {
	StatusPageId int64 `json:"status_page_id" validate:"required"`
	MonitorId    int64 `json:"monitor_id" validate:"required"`
}
