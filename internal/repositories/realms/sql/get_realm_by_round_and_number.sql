SELECT id, round_id, number, name
FROM realms
WHERE round_id = $1 AND number = $2;
