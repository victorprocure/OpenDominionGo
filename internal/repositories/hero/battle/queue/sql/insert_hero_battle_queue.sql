INSERT INTO hero_battle_queue (hero_id, level, rating)
VALUES ($1, $2, $3)
RETURNING id;
