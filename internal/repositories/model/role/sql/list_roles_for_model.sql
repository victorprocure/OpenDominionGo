SELECT role_id
FROM model_has_roles
WHERE model_type = $1 AND model_id = $2
ORDER BY role_id ASC
LIMIT $3 OFFSET $4;
