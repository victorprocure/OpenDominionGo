SELECT id, message_board_thread_id, user_id, body, flagged_for_removal, flagged_by
FROM message_board_posts
WHERE message_board_thread_id = $1
ORDER BY id ASC;
