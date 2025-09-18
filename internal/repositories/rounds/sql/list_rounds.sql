SELECT id, number, name
FROM rounds
ORDER BY number DESC
LIMIT $1 OFFSET $2;
