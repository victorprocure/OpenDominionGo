INSERT INTO dominion_techs (dominion_id, tech_id)
VALUES ($1, $2)
RETURNING id;
