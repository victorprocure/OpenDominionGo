DELETE FROM role_has_permissions
WHERE role_id = $1 AND permission_id = $2;
