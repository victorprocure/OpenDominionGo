INSERT INTO realm_wars (source_realm_id, target_realm_id, active_at)
VALUES ($1, $2, COALESCE($3, NOW()))
RETURNING id;
