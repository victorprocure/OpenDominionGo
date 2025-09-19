SELECT id, source_realm_id, target_realm_id, active_at, inactive_at
FROM realm_wars
WHERE inactive_at IS NULL AND (source_realm_id = $1 OR target_realm_id = $1)
ORDER BY active_at DESC NULLS LAST, id DESC;
