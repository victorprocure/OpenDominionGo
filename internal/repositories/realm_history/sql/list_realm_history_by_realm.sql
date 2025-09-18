SELECT id, realm_id, dominion_id, event, delta, created_at
FROM realm_history
WHERE realm_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
