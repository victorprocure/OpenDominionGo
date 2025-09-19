SELECT id, key
FROM tech_perk_types
ORDER BY id DESC
LIMIT $1 OFFSET $2;
