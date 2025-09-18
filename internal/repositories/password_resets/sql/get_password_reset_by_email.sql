SELECT email, token, created_at
FROM password_resets
WHERE email = $1;
