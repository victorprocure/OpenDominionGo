SELECT id, realm_id, dominion_id, title, deleted_at, last_activity
FROM council_threads
WHERE realm_id = $1
ORDER BY last_activity DESC NULLS LAST, id DESC
LIMIT $2 OFFSET $3;
