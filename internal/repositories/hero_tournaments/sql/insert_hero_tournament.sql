INSERT INTO hero_tournaments (round_id, name, start_date)
VALUES ($1, $2, $3)
RETURNING id;
