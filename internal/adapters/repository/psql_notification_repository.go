package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/queries"
)

type PsqlNotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *PsqlNotificationRepository {
	return &PsqlNotificationRepository{
		db: db,
	}
}

func (p *PsqlNotificationRepository) Save(ctx context.Context, n *domain.Notification) (*domain.Notification, error) {
	stmt, _ := p.db.PrepareContext(ctx, queries.NotificationCreateQuery)

	defer stmt.Close()

	var id int64

	jsonProviderData, _ := json.Marshal(n.ProviderData)

	err := stmt.QueryRowContext(
		ctx,
		n.Title,
		n.Description,
		n.Provider,
		jsonProviderData,
	).Scan(&id)

	n.Id = id

	if err != nil {
		return nil, err
	}

	return n, nil
}

func (p *PsqlNotificationRepository) AssignToMonitor(ctx context.Context, notificationId int64, monitorId int64) error {
	stmt, _ := p.db.PrepareContext(ctx, queries.NotificationAssignToMonitorQuery)

	log.Printf("Assigning notification %d to monitor %d\n\n", notificationId, monitorId)

	defer stmt.Close()

	_, err := stmt.ExecContext(ctx, monitorId, notificationId)

	if err != nil {
		return err
	}

	return nil
}

func (p *PsqlNotificationRepository) UnassignToMonitor(ctx context.Context, notificationId int64, monitorId int64) error {
	stmt, _ := p.db.PrepareContext(ctx, queries.NotificationUnassignToMonitorQuery)

	log.Printf("Unassigning notification %d to monitor %d\n\n", notificationId, monitorId)

	defer stmt.Close()

	_, err := stmt.ExecContext(ctx, monitorId, notificationId)

	if err != nil {
		return err
	}

	return nil
}

func (p *PsqlNotificationRepository) GetById(ctx context.Context, id int64) (*domain.Notification, error) {
	stmt, _ := p.db.PrepareContext(ctx, queries.NotificationGetByIdQuery)

	defer stmt.Close()

	var n domain.Notification
	var jsonProvider string

	err := stmt.QueryRowContext(ctx, id).Scan(
		&n.Id,
		&n.Title,
		&n.Description,
		&n.Provider,
		&jsonProvider,
		&n.CreatedAt,
		&n.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonProvider), &n.ProviderData)

	if err != nil {
		return nil, err
	}

	return &n, nil
}

func (p *PsqlNotificationRepository) GetByMonitorId(ctx context.Context, monitorId int64) ([]*domain.Notification, error) {
	stmt, _ := p.db.PrepareContext(ctx, queries.NotificationGetByMonitorIdQuery)

	defer stmt.Close()

	var notifications []*domain.Notification

	rows, err := stmt.QueryContext(ctx, monitorId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var n domain.Notification
		var jsonProvider string

		err = rows.Scan(
			&n.Id,
			&n.Title,
			&n.Description,
			&n.Provider,
			&jsonProvider,
			&n.CreatedAt,
			&n.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(jsonProvider), &n.ProviderData)

		if err != nil {
			return nil, err
		}

		notifications = append(notifications, &n)
	}

	return notifications, nil
}

func (p *PsqlNotificationRepository) Update(ctx context.Context, n *domain.Notification) (*domain.Notification, error) {
	stmt, _ := p.db.PrepareContext(ctx, queries.NotificationUpdateQuery)

	defer stmt.Close()
	jsonProviderData, _ := json.Marshal(n.ProviderData)

	_, err := stmt.ExecContext(
		ctx,
		n.Title,
		n.Description,
		n.Provider,
		jsonProviderData)

	n.UpdatedAt = time.Now()

	if err != nil {
		return nil, err
	}

	return n, nil
}

func (p *PsqlNotificationRepository) Delete(ctx context.Context, id int64) (int64, error) {
	stmt, _ := p.db.PrepareContext(ctx, queries.NotificationDeleteQuery)

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
