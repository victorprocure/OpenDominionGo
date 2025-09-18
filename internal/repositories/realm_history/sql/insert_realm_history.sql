INSERT INTO realm_history (realm_id, dominion_id, event, delta)
VALUES ($1, $2, $3, $4)
RETURNING id;
