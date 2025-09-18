INSERT INTO roles (name, guard_name)
VALUES ($1, $2)
RETURNING id;
