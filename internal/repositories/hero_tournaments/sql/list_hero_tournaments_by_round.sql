SELECT id, round_id, name, current_round_number, finished, winner_dominion_id, start_date
FROM hero_tournaments
WHERE round_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
