SELECT entry_uuid, tag
FROM telescope_entries_tags
WHERE entry_uuid = $1
ORDER BY tag
LIMIT $2 OFFSET $3;
