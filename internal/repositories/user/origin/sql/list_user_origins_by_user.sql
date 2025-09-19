SELECT id, user_id, dominion_id, ip_address, count
FROM user_origins
WHERE user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
