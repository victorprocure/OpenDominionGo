INSERT INTO user_origins (user_id, dominion_id, ip_address, count)
VALUES ($1, $2, $3, $4)
RETURNING id;
