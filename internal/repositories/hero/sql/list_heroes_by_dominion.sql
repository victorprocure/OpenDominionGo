SELECT id, dominion_id, name, class, experience, combat_rating
FROM heroes
WHERE dominion_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
