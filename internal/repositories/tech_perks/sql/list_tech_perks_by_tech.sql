SELECT id, tech_id, tech_perk_type_id, value
FROM tech_perks
WHERE tech_id = $1
ORDER BY id ASC
LIMIT $2 OFFSET $3;
