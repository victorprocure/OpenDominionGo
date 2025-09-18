INSERT INTO council_threads (realm_id, dominion_id, title, body)
VALUES ($1, $2, $3, $4)
RETURNING id;
