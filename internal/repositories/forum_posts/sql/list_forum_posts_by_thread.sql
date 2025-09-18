SELECT id, forum_thread_id, dominion_id, body, flagged_for_removal, flagged_by
FROM forum_posts
WHERE forum_thread_id = $1
ORDER BY id ASC
LIMIT $2 OFFSET $3;
