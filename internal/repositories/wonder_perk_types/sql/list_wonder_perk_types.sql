SELECT id, key
FROM wonder_perk_types
ORDER BY id ASC
LIMIT $1 OFFSET $2;
