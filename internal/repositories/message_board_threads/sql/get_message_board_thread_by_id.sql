SELECT id, message_board_category_id, user_id, title, body, flagged_for_removal, flagged_by, last_activity
FROM message_board_threads
WHERE id = $1;
