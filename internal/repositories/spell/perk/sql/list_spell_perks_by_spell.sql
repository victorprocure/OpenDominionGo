SELECT id, spell_id, spell_perk_type_id, value
FROM spell_perks
WHERE spell_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
