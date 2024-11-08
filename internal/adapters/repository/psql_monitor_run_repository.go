package repository

import (
	"context"
	"database/sql"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/queries"
)

type PsqlMonitorRunRepository struct {
	db *sql.DB
}

func NewMonitorRunRepository(db *sql.DB) *PsqlMonitorRunRepository {
	return &PsqlMonitorRunRepository{
		db: db,
	}
}

func (p *PsqlMonitorRunRepository) Save(ctx context.Context, mr *domain.MonitorRun) (*domain.MonitorRun, error) {
	stmt, err := p.db.PrepareContext(ctx, queries.MonitorRunCreateQuery)

	if err != nil {
		return nil, err
	}
	var id int64

	defer stmt.Close()

	err = stmt.QueryRowContext(
		ctx,
		mr.MonitorId,
		mr.Status,
		mr.DnsStartedAt,
		mr.DnsEndedAt,
		mr.TlsHandshakeStartedAt,
		mr.TlsHandshakeEndedAt,
		mr.ConnectStartedAt,
		mr.ConnectEndedAt,
		mr.Latency,
		mr.RanAt).Scan(&id)

	mr.Id = id
	return mr, nil
}

func (p *PsqlMonitorRunRepository) GetById(ctx context.Context, id int64) (*domain.MonitorRun, error) {
	return nil, nil
}

func (p *PsqlMonitorRunRepository) GetMonitorRunsTimeSeries90d(ctx context.Context, monitorId int64) ([]*domain.UptimeEntry, error) {
	stmt, err := p.db.PrepareContext(ctx, queries.MonitorRunGetTimeSeries90dQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, monitorId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var uptimeEntries []*domain.UptimeEntry

	for rows.Next() {
		var uptimeEntry domain.UptimeEntry

		err = rows.Scan(
			&uptimeEntry.Key,
			&uptimeEntry.Ok,
			&uptimeEntry.Total,
			&uptimeEntry.Uptime,
		)

		if err != nil {
			return nil, err
		}

		uptimeEntries = append(uptimeEntries, &uptimeEntry)
	}

	return uptimeEntries, nil
}
