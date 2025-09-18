INSERT INTO round_wonders (round_id, realm_id, wonder_id, power)
VALUES ($1,$2,$3,$4)
RETURNING id;
