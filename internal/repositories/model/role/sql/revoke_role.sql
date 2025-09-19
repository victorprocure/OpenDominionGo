DELETE FROM model_has_roles
WHERE role_id = $1 AND model_type = $2 AND model_id = $3;
