UPDATE rounds
SET name = $2,
    updated_at = now()
WHERE id = $1;
