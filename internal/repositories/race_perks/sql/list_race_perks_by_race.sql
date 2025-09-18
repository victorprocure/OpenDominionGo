SELECT id, race_id, race_perk_type_id, value
FROM race_perks
WHERE race_id = $1
ORDER BY id ASC
LIMIT $2 OFFSET $3;
