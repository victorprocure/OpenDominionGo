INSERT INTO hero_battles (round_id, current_turn, pvp)
VALUES ($1, $2, $3)
RETURNING id;
