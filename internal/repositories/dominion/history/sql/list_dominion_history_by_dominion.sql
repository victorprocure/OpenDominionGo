SELECT id, dominion_id, event, delta, created_at, ip, device
FROM dominion_history
WHERE dominion_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
