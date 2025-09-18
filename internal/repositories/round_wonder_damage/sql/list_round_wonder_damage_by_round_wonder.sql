SELECT id, round_wonder_id, realm_id, dominion_id, damage, source
FROM round_wonder_damage
WHERE round_wonder_id = $1
ORDER BY id ASC;
