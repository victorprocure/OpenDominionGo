INSERT INTO spell_perk_types (key)
VALUES ($1)
RETURNING id;
