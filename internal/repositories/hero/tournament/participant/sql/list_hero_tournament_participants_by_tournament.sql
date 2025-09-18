SELECT id, hero_tournament_id, hero_id, wins, losses, draws, standing, eliminated
FROM hero_tournament_participants
WHERE hero_tournament_id = $1
ORDER BY standing NULLS LAST, id ASC
LIMIT $2 OFFSET $3;
