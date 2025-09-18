SELECT id, council_thread_id, dominion_id, body, deleted_at
FROM council_posts
WHERE council_thread_id = $1
ORDER BY id ASC
LIMIT $2 OFFSET $3;
