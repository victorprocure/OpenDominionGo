INSERT INTO message_board_posts (message_board_thread_id, user_id, body)
VALUES ($1, $2, $3)
RETURNING id;
