INSERT INTO dominion_spells (dominion_id, duration, cast_by_dominion_id, spell_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT (dominion_id, spell_id) DO UPDATE
SET duration = EXCLUDED.duration,
    cast_by_dominion_id = EXCLUDED.cast_by_dominion_id,
    updated_at = NOW();
