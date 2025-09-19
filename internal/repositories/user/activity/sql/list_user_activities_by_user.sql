SELECT id, user_id, ip, key, context, status, device
FROM user_activities
WHERE user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
