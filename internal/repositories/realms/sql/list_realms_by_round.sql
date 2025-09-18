SELECT id, round_id, number, name
FROM realms
WHERE round_id = $1
ORDER BY number ASC
LIMIT $2 OFFSET $3;
