package inputs

type AssignMonitorToNotificationInput struct {
	NotificationId int64 `json:"notification_id" validate:"required"`
	MonitorId      int64 `json:"monitor_id" validate:"required"`
}

type UnassignMonitorFromNotificationInput struct {
	NotificationId int64 `json:"notification_id" validate:"required"`
	MonitorId      int64 `json:"monitor_id" validate:"required"`
}
