SELECT id, round_id, realm_id, race_id, name
FROM dominions
WHERE realm_id = $1
ORDER BY id ASC;
