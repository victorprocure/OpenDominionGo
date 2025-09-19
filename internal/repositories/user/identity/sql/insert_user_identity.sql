INSERT INTO user_identities (user_id, fingerprint, user_agent, count)
VALUES ($1, $2, $3, $4)
RETURNING id;
