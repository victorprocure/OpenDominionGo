SELECT id, source_id, target_id, endorsed
FROM user_feedback
WHERE target_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
