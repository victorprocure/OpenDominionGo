INSERT INTO race_perks (race_id, race_perk_type_id, value)
VALUES ($1, $2, $3)
RETURNING id;
