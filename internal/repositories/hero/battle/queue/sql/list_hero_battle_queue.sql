SELECT id, hero_id, level, rating
FROM hero_battle_queue
ORDER BY id DESC
LIMIT $1 OFFSET $2;
