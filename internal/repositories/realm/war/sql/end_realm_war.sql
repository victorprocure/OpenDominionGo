UPDATE realm_wars
SET inactive_at = NOW()
WHERE id = $1;
