INSERT INTO user_feedback (source_id, target_id, endorsed)
VALUES ($1, $2, $3)
RETURNING id;
