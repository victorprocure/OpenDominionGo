SELECT id, name, slug, role_required
FROM message_board_categories
WHERE slug = $1;
