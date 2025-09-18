SELECT id, user_id, achievement_id
FROM user_achievements
WHERE user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
