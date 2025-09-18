SELECT id, dominion_id, tech_id
FROM dominion_techs
WHERE dominion_id = $1
ORDER BY id DESC;
