SELECT id, name, description, icon
FROM achievements
ORDER BY id DESC
LIMIT $1 OFFSET $2;
