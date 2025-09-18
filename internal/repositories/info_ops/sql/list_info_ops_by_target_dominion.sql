SELECT id, source_realm_id, source_dominion_id, target_dominion_id, target_realm_id, type, data
FROM info_ops
WHERE target_dominion_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
