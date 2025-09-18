INSERT INTO game_events (round_id, source_type, source_id, target_type, target_id, type, data)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;
