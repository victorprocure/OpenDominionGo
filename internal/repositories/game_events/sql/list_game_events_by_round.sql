SELECT id, round_id, source_type, source_id, target_type, target_id, type, data
FROM game_events
WHERE round_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
