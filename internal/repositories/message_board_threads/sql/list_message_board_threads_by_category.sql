SELECT id, message_board_category_id, user_id, title, flagged_for_removal, flagged_by, last_activity
FROM message_board_threads
WHERE message_board_category_id = $1
ORDER BY last_activity DESC NULLS LAST, id DESC
LIMIT $2 OFFSET $3;
