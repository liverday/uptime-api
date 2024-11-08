package domain

import "context"

type StatusPageRepository interface {
	Save(ctx context.Context, sp *StatusPage) (*StatusPage, error)
	GetById(ctx context.Context, id int64) (*StatusPage, error)
	AssignMonitors(ctx context.Context, pageId int64, monitorIds []int64) error
	UnassignMonitors(ctx context.Context, pageId int64, monitorIds []int64) error
}
