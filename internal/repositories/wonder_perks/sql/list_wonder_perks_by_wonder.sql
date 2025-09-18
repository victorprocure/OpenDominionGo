SELECT id, wonder_id, wonder_perk_type_id, value
FROM wonder_perks
WHERE wonder_id = $1
ORDER BY id ASC
LIMIT $2 OFFSET $3;
