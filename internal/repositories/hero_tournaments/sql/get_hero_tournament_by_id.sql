SELECT id, round_id, name, current_round_number, finished, winner_dominion_id, start_date
FROM hero_tournaments
WHERE id = $1;
