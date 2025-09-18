UPDATE achievements
SET name = $2,
    description = $3,
    icon = $4,
    updated_at = NOW()
WHERE id = $1;
