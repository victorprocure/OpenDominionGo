SELECT id, name, guard_name
FROM permissions
WHERE name = $1;
