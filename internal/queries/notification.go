package queries

const (
	NotificationCreateQuery = `
		INSERT INTO notifications (title, description, provider, provider_data)
		VALUES ($1, $2, $3, $4) RETURNING id
	`

	NotificationAssignToMonitorQuery = `
		INSERT INTO monitor_notifications (monitor_id, notification_id) VALUES ($1, $2) ON CONFLICT DO NOTHING
	`

	NotificationUnassignToMonitorQuery = `
		DELETE FROM monitor_notifications WHERE monitor_id = $1 AND notification_id = $2
	`

	NotificationGetByMonitorIdQuery = `
		SELECT 
			n.id, 
			n.title, 
			n.description, 
			n.provider, 
			n.provider_data,
			n.created_at,
			n.updated_at
		FROM notifications n INNER JOIN public.monitor_notifications mn on n.id = mn.notification_id
		WHERE mn.monitor_id = $1
	`

	NotificationGetByIdQuery = `
		SELECT 
			id, 
			title, 
			description, 
			provider, 
			provider_data,
			created_at,
			updated_at
		FROM notifications WHERE id = $1
	`

	NotificationUpdateQuery = `
		UPDATE notifications SET 
			 title = $1, 
			 description = $2, 
			 provider = $3, 
			 provider_data = $4,
			 updated_at = now()
		 WHERE id = $5
	`

	NotificationDeleteQuery = `
		DELETE FROM monitor_notifications WHERE notification_id = $1;
		DELETE FROM notifications WHERE id = $1;
	`
)
