package repository

import (
	"context"
	"database/sql"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/queries"
)

type PsqlStatusPageRepository struct {
	db *sql.DB
}

func NewStatusPageRepository(db *sql.DB) *PsqlStatusPageRepository {
	return &PsqlStatusPageRepository{
		db: db,
	}
}

func (p *PsqlStatusPageRepository) Save(ctx context.Context, sp *domain.StatusPage) (*domain.StatusPage, error) {
	stmt, err := p.db.PrepareContext(ctx, queries.StatusPageCreateQuery)

	if err != nil {
		return nil, err
	}

	var id int64

	err = stmt.QueryRowContext(ctx, sp.Title, sp.Description).Scan(&id)

	if err != nil {
		return nil, err
	}

	sp.Id = id
	return sp, nil
}

func (p *PsqlStatusPageRepository) GetById(ctx context.Context, id int64) (*domain.StatusPage, error) {
	stmt, err := p.db.PrepareContext(ctx, queries.StatusPageGetByIdQuery)

	if err != nil {
		return nil, err
	}

	var sp domain.StatusPage

	err = stmt.QueryRowContext(ctx, id).Scan(
		&sp.Id,
		&sp.Title,
		&sp.Description,
		&sp.CreatedAt,
		&sp.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &sp, nil
}

func (p *PsqlStatusPageRepository) AssignMonitors(ctx context.Context, pageId int64, monitorIds []int64) error {
	tx, err := p.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, queries.StatusPageAssignMonitorsQuery)

	if err != nil {
		return err
	}

	for _, monitorId := range monitorIds {
		_, err = stmt.ExecContext(ctx, monitorId, pageId)

		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *PsqlStatusPageRepository) UnassignMonitors(ctx context.Context, pageId int64, monitorIds []int64) error {
	tx, err := p.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, queries.StatusPageUnassignMonitorsQuery)

	if err != nil {
		return err
	}

	for _, monitorId := range monitorIds {
		_, err = stmt.ExecContext(ctx, monitorId, pageId)

		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
