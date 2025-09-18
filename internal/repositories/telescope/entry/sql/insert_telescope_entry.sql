INSERT INTO telescope_entries (uuid, batch_id, family_hash, should_display_on_index, type, content)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING sequence;
