SELECT tag
FROM telescope_monitoring
ORDER BY tag ASC
LIMIT $1 OFFSET $2;
