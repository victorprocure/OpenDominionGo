SELECT dominion_id, duration, cast_by_dominion_id, spell_id
FROM dominion_spells
WHERE dominion_id = $1
ORDER BY spell_id;
