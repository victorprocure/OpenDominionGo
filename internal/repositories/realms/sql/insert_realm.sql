INSERT INTO realms (round_id, number, name, alignment)
VALUES ($1,$2,$3,$4)
RETURNING id;
