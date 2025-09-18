INSERT INTO message_board_threads (message_board_category_id, user_id, title, body)
VALUES ($1, $2, $3, $4)
RETURNING id;
