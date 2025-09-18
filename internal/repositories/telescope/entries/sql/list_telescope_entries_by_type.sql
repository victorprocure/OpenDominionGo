SELECT sequence, uuid, batch_id, family_hash, should_display_on_index, type, content
FROM telescope_entries
WHERE type = $1
ORDER BY sequence DESC
LIMIT $2 OFFSET $3;
