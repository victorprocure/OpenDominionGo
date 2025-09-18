INSERT INTO heroes (dominion_id, name, class)
VALUES ($1, $2, $3)
RETURNING id;
