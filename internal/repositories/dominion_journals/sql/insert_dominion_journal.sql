INSERT INTO dominion_journals (dominion_id, content)
VALUES ($1, $2)
RETURNING id;
