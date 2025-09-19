SELECT id, realm_id, dominion_id, title, body, deleted_at, last_activity
FROM council_threads
WHERE id = $1;
