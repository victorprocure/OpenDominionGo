SELECT sequence, uuid, batch_id, family_hash, should_display_on_index, type, content
FROM telescope_entries
WHERE sequence = $1;
