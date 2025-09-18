INSERT INTO unit_perks (race_id, unit_slot, unit_perk_type_id, value)
VALUES ($1, $2, $3, $4)
RETURNING id;
