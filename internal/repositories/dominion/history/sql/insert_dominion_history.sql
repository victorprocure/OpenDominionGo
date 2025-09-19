INSERT INTO dominion_history (dominion_id, event, delta, ip, device)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
