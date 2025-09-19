INSERT INTO council_posts (council_thread_id, dominion_id, body)
VALUES ($1, $2, $3)
RETURNING id;
