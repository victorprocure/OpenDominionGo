SELECT id, user_id, fingerprint, user_agent, count
FROM user_identities
WHERE user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
