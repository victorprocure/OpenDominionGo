UPDATE message_board_categories
SET name = $2,
    slug = $3,
    role_required = $4
WHERE id = $1;
