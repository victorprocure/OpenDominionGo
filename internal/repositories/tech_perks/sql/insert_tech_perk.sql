INSERT INTO tech_perks (tech_id, tech_perk_type_id, value)
VALUES ($1, $2, $3)
RETURNING id;
