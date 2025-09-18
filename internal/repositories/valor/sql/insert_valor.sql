INSERT INTO valor (round_id, realm_id, dominion_id, source, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
