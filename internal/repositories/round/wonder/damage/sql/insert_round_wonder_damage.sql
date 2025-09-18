INSERT INTO round_wonder_damage (round_wonder_id, realm_id, dominion_id, damage, source)
VALUES ($1,$2,$3,$4,$5)
RETURNING id;
