package queries

const (
	StatusPageCreateQuery = `
		INSERT INTO status_pages (title, description) VALUES ($1, $2) RETURNING id
	`

	StatusPageGetByIdQuery = `
		SELECT 
		    id,
		    title,
		    description,
		    created_at,
			updated_at
		FROM status_pages WHERE id = $1
	`

	StatusPageAssignMonitorsQuery = `
		INSERT INTO monitor_status_pages (monitor_id, status_page_id) VALUES ($1, $2) ON CONFLICT DO NOTHING
	`

	StatusPageUnassignMonitorsQuery = `
		DELETE FROM monitor_status_pages WHERE monitor_id = $1 AND status_page_id = $2
	`
)
