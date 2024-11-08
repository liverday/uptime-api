package domain

import "context"

type NotificationRepository interface {
	Save(ctx context.Context, n *Notification) (*Notification, error)
	AssignToMonitor(ctx context.Context, notificationId int64, monitorId int64) error
	GetById(ctx context.Context, id int64) (*Notification, error)
	GetByMonitorId(ctx context.Context, monitorId int64) ([]*Notification, error)
	Update(ctx context.Context, n *Notification) (*Notification, error)
	Delete(ctx context.Context, id int64) (int64, error)
	UnassignToMonitor(ctx context.Context, notificationId int64, monitorId int64) error
}
