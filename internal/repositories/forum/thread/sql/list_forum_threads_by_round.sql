SELECT id, round_id, dominion_id, title, flagged_for_removal, flagged_by, last_activity
FROM forum_threads
WHERE round_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
