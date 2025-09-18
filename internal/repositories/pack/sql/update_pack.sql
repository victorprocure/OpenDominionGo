UPDATE packs
SET realm_id = $2,
    name = $3,
    password = $4,
    size = $5,
    updated_at = NOW()
WHERE id = $1;
