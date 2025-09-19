SELECT id, user_id, discord_user_id, username, discriminator, email, refresh_token
FROM user_discord_users
WHERE user_id = $1;
