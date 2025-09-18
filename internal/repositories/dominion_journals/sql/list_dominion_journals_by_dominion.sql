SELECT id, dominion_id, content, created_at, updated_at
FROM dominion_journals
WHERE dominion_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
