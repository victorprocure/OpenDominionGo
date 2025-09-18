INSERT INTO round_leagues (key, description)
VALUES ($1, $2)
ON CONFLICT (key) DO UPDATE
SET description = EXCLUDED.description,
    updated_at = NOW()
RETURNING id;
