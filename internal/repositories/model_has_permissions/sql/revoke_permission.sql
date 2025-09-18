DELETE FROM model_has_permissions
WHERE permission_id = $1 AND model_type = $2 AND model_id = $3;
