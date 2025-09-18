INSERT INTO permissions (name, guard_name)
VALUES ($1, $2)
RETURNING id;
