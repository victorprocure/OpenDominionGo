SELECT id, connection, queue, payload, exception
FROM failed_jobs
ORDER BY id DESC
LIMIT $1 OFFSET $2;
