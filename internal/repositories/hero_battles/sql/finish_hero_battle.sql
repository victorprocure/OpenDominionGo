UPDATE hero_battles
SET finished = true,
    winner_combatant_id = $2,
    last_processed_at = NOW()
WHERE id = $1;
