SELECT dominion_id, source, resource, hours, amount
FROM dominion_queue
WHERE dominion_id = $1
ORDER BY source, resource;
