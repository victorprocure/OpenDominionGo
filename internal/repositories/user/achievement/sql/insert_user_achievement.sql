INSERT INTO user_achievements (user_id, achievement_id)
VALUES ($1, $2)
RETURNING id;
