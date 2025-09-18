SELECT id, race_id, unit_slot, unit_perk_type_id, value
FROM unit_perks
WHERE race_id = $1 AND unit_slot = $2
ORDER BY id ASC
LIMIT $3 OFFSET $4;
