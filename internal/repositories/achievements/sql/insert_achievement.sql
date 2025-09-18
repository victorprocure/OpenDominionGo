INSERT INTO achievements (name, description, icon)
VALUES ($1, $2, $3)
RETURNING id;
