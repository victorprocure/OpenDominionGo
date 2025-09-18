SELECT id, name, description, icon
FROM achievements
WHERE name = $1;
