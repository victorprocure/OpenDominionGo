INSERT INTO jobs (queue, payload, attempts, available_at, created_at)
VALUES ($1, $2, 0, EXTRACT(EPOCH FROM NOW())::int, EXTRACT(EPOCH FROM NOW())::int)
RETURNING id;
