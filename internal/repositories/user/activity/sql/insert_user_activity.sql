INSERT INTO user_activities (user_id, ip, key, context, status, device)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
