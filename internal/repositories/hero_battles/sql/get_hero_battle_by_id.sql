SELECT id, round_id, current_turn, winner_combatant_id, finished, last_processed_at, pvp
FROM hero_battles
WHERE id = $1;
