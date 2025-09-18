SELECT id, round_id, realm_id, wonder_id, power
FROM round_wonders
WHERE round_id = $1
ORDER BY id ASC;
