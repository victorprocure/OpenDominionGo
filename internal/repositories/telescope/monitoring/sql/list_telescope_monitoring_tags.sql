SELECT tag
FROM telescope_monitoring
ORDER BY tag
LIMIT $1 OFFSET $2;
