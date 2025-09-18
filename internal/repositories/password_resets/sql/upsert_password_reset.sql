INSERT INTO password_resets (email, token, created_at)
VALUES ($1, $2, NOW())
ON CONFLICT (email) DO UPDATE
SET token = EXCLUDED.token,
    created_at = NOW();
