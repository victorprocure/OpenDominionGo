INSERT INTO wonder_perks (wonder_id, wonder_perk_type_id, value)
VALUES ($1, $2, $3)
RETURNING id;
