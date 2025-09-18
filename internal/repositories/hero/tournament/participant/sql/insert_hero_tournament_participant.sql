INSERT INTO hero_tournament_participants (hero_tournament_id, hero_id)
VALUES ($1, $2)
RETURNING id;
