INSERT INTO message_board_categories (name, slug, role_required)
VALUES ($1, $2, $3)
RETURNING id;
