SELECT id, round_id, realm_id, dominion_id, source, amount
FROM valor
WHERE dominion_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
