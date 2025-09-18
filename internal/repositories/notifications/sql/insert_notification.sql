INSERT INTO notifications (type, notifiable_type, notifiable_id, data)
VALUES ($1, $2, $3, $4)
RETURNING id;
