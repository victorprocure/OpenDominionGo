SELECT id, round_id, dominion_id, title, body, flagged_for_removal, flagged_by, last_activity
FROM forum_threads
WHERE id = $1;
