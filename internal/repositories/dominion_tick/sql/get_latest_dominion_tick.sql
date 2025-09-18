SELECT id, dominion_id, prestige, peasants, morale, updated_at
FROM dominion_tick
WHERE dominion_id = $1
ORDER BY id DESC
LIMIT 1;
