SELECT id, round_id, realm_id, race_id, name
FROM dominions
WHERE round_id = $1 AND name = $2;
