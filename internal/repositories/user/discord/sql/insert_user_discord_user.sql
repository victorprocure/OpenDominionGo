INSERT INTO user_discord_users (user_id, discord_user_id, username, discriminator, email, refresh_token, expires_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;
