UPDATE dominion_spells
SET duration = $2, cast_by_dominion_id = $3
WHERE dominion_id = $1 AND spell_id = $4;

INSERT INTO dominion_spells (dominion_id, duration, cast_by_dominion_id, spell_id)
SELECT $1, $2, $3, $4
WHERE NOT EXISTS (
    SELECT 1 FROM dominion_spells WHERE dominion_id = $1 AND spell_id = $4
);
