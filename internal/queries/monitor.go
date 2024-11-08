package queries

const (
	MonitorCreateQuery = `INSERT INTO monitors (
			title,
			description,
			status,
			periodicity,
			url,
			method,
			headers,
			body,
			degraded_after) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	MonitorSelectByIdQuery = `
		SELECT 
		    id,
		    title,
		    description,
		    status,
		    periodicity,
		    url,
		    method,
		    headers,
		    body,
		    degraded_after,
		    created_at,
		    updated_at    
		FROM monitors WHERE id = $1
	`

	MonitorSelectAllByIdsQuery = `
		SELECT 
		    id,
		    title,
		    description,
		    status,
		    periodicity,
		    url,
		    method,
		    headers,
		    body,
		    degraded_after,
		    created_at,
		    updated_at    
		FROM monitors WHERE id = ANY($1::int[])
	`

	MonitorSelectByPeriodicityQuery = `
		SELECT 
		    id,
		    title,
		    description,
		    status,
		    periodicity,
		    url,
		    method,
		    headers,
		    body,
		    degraded_after,
		    created_at,
		    updated_at 
		FROM monitors WHERE periodicity = $1
	`

	MonitorUpdateQuery = `
		UPDATE monitors SET
			title = $1,
			description = $2,
			status = $3,
			periodicity = $4,
			url = $5,
			method = $6,
			headers = $7,
			body = $8,
			degraded_after = $9,
			updated_at = now()
		WHERE id = $10
	`

	MonitorUpdateStatusQuery = `
		UPDATE monitors SET
			status = $1,
			updated_at = now()
		WHERE id = $2
	`

	MonitorDeleteQuery = `
		DELETE FROM monitors WHERE id = $1
	`

	MonitorSelectAllByStatusPageId = ` 
		SELECT 
		    id,
		    title,
		    description,
		    status,
		    periodicity,
		    url,
		    method,
		    headers,
		    body,
		    degraded_after,
		    created_at,
		    updated_at  
		FROM monitors
		WHERE id in (SELECT monitor_id FROM monitor_status_pages WHERE status_page_id = $1)
		ORDER BY id
	`
)
