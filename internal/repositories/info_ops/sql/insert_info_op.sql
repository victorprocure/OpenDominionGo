INSERT INTO info_ops (source_realm_id, source_dominion_id, target_dominion_id, type, data, target_realm_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
