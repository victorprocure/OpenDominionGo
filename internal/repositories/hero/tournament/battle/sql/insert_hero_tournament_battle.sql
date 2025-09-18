INSERT INTO hero_tournament_battles (hero_tournament_id, hero_battle_id, round_number)
VALUES ($1, $2, $3)
RETURNING id;
