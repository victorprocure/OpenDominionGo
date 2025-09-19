INSERT INTO forum_posts (forum_thread_id, dominion_id, body)
VALUES ($1, $2, $3)
RETURNING id;
