INSERT INTO spell_perks (spell_id, spell_perk_type_id, value)
VALUES ($1, $2, $3)
RETURNING id;
