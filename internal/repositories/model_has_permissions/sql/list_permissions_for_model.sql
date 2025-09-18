SELECT permission_id
FROM model_has_permissions
WHERE model_type = $1 AND model_id = $2
ORDER BY permission_id ASC
LIMIT $3 OFFSET $4;
