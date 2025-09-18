SELECT id, round_id, realm_id, name, size
FROM packs
WHERE round_id = $1 AND name = $2;
