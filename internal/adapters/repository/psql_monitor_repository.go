package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"time"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/queries"
)

type PsqlMonitorRepository struct {
	db *sql.DB
}

func NewMonitorRepository(db *sql.DB) *PsqlMonitorRepository {
	return &PsqlMonitorRepository{
		db: db,
	}
}

func (p *PsqlMonitorRepository) Save(ctx context.Context, m *domain.Monitor) (*domain.Monitor, error) {
	stmt, err := p.db.PrepareContext(ctx, queries.MonitorCreateQuery)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var id int64

	err = stmt.QueryRowContext(
		ctx,
		m.Title,
		m.Description,
		m.Status,
		m.Periodicity,
		m.Url,
		m.Method,
		m.Headers,
		m.Body,
		m.DegradedAfter).Scan(&id)

	if err != nil {
		return nil, err
	}

	m.Id = id
	return m, nil
}

func (p *PsqlMonitorRepository) GetById(ctx context.Context, id int64) (*domain.Monitor, error) {
	stmt, _ := p.db.PrepareContext(ctx, queries.MonitorSelectByIdQuery)

	defer stmt.Close()

	var m domain.Monitor

	err := stmt.QueryRowContext(ctx, id).Scan(
		&m.Id,
		&m.Title,
		&m.Description,
		&m.Status,
		&m.Periodicity,
		&m.Url,
		&m.Method,
		&m.Headers,
		&m.Body,
		&m.DegradedAfter,
		&m.CreatedAt,
		&m.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (p *PsqlMonitorRepository) GetAllByIds(ctx context.Context, ids []int64) ([]*domain.Monitor, error) {
	if len(ids) == 0 {
		return nil, errors.New("ids is required")
	}

	stmt, err := p.db.PrepareContext(ctx, queries.MonitorSelectAllByIdsQuery)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, pq.Array(ids))

	if err != nil {
		return nil, err
	}

	var monitors []*domain.Monitor

	for rows.Next() {
		var m domain.Monitor

		err := rows.Scan(
			&m.Id,
			&m.Title,
			&m.Description,
			&m.Status,
			&m.Periodicity,
			&m.Url,
			&m.Method,
			&m.Headers,
			&m.Body,
			&m.DegradedAfter,
			&m.CreatedAt,
			&m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		monitors = append(monitors, &m)
	}

	return monitors, nil
}

func (p *PsqlMonitorRepository) GetAllByStatusPageId(ctx context.Context, statusPageId int64) ([]*domain.Monitor, error) {
	stmt, err := p.db.PrepareContext(ctx, queries.MonitorSelectAllByStatusPageId)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var monitors []*domain.Monitor

	rows, err := stmt.QueryContext(ctx, statusPageId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var m domain.Monitor

		err := rows.Scan(
			&m.Id,
			&m.Title,
			&m.Description,
			&m.Status,
			&m.Periodicity,
			&m.Url,
			&m.Method,
			&m.Headers,
			&m.Body,
			&m.DegradedAfter,
			&m.CreatedAt,
			&m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		monitors = append(monitors, &m)
	}

	return monitors, nil
}

func (p *PsqlMonitorRepository) GetAllByPeriodicity(ctx context.Context, pd domain.Periodicity) ([]*domain.Monitor, error) {
	stmt, err := p.db.PrepareContext(ctx, queries.MonitorSelectByPeriodicityQuery)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, pd)

	if err != nil {
		return nil, err
	}

	var monitors []*domain.Monitor

	for rows.Next() {
		var m domain.Monitor

		err := rows.Scan(
			&m.Id,
			&m.Title,
			&m.Description,
			&m.Status,
			&m.Periodicity,
			&m.Url,
			&m.Method,
			&m.Headers,
			&m.Body,
			&m.DegradedAfter,
			&m.CreatedAt,
			&m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		monitors = append(monitors, &m)
	}

	return monitors, nil
}

func (p *PsqlMonitorRepository) Update(ctx context.Context, m *domain.Monitor) (*domain.Monitor, error) {
	stmt, _ := p.db.PrepareContext(ctx, queries.MonitorUpdateQuery)

	defer stmt.Close()

	_, err := stmt.ExecContext(
		ctx,
		m.Title,
		m.Description,
		m.Status,
		m.Periodicity,
		m.Url,
		m.Method,
		m.Headers,
		m.Body,
		m.DegradedAfter,
		m.Id)

	m.UpdatedAt = time.Now().UTC()

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (p *PsqlMonitorRepository) UpdateStatus(ctx context.Context, monitorId int64, status domain.MonitorStatus) error {
	stmt, _ := p.db.PrepareContext(ctx, queries.MonitorUpdateStatusQuery)

	defer stmt.Close()

	_, err := stmt.ExecContext(ctx, status, monitorId)

	if err != nil {
		return err
	}

	return nil
}

func (p *PsqlMonitorRepository) Delete(ctx context.Context, id int64) (int64, error) {
	stmt, _ := p.db.PrepareContext(ctx, queries.MonitorDeleteQuery)

	defer stmt.Close()

	exec, err := stmt.ExecContext(ctx, id)

	if err != nil {
		return 0, err
	}

	return exec.RowsAffected()
}
