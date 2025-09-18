SELECT id, type, notifiable_type, notifiable_id, data
FROM notifications
WHERE notifiable_type = $1 AND notifiable_id = $2
ORDER BY id DESC
LIMIT $3 OFFSET $4;
