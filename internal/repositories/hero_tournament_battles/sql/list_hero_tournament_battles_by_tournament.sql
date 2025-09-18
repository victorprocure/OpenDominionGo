SELECT id, hero_tournament_id, hero_battle_id, round_number
FROM hero_tournament_battles
WHERE hero_tournament_id = $1
ORDER BY round_number ASC, id ASC
LIMIT $2 OFFSET $3;
