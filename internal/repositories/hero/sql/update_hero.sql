UPDATE heroes
SET name = COALESCE($2, name),
    class = COALESCE($3, class),
    experience = COALESCE($4, experience)
WHERE id = $1;
