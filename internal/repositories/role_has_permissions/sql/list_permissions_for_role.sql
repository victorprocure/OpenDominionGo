SELECT permission_id
FROM role_has_permissions
WHERE role_id = $1
ORDER BY permission_id ASC
LIMIT $2 OFFSET $3;
