SELECT id, key
FROM race_perk_types
ORDER BY id ASC
LIMIT $1 OFFSET $2;
