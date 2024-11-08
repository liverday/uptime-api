package domain

import "context"

type MonitorRepository interface {
	Save(context context.Context, m *Monitor) (*Monitor, error)
	GetById(context context.Context, id int64) (*Monitor, error)
	GetAllByIds(context context.Context, ids []int64) ([]*Monitor, error)
	GetAllByStatusPageId(context context.Context, statusPageId int64) ([]*Monitor, error)
	GetAllByPeriodicity(context context.Context, pd Periodicity) ([]*Monitor, error)
	Delete(context context.Context, id int64) (int64, error)
	Update(context context.Context, m *Monitor) (*Monitor, error)
	UpdateStatus(context context.Context, monitorId int64, status MonitorStatus) error
}
