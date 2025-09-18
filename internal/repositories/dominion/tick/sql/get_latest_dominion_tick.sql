SELECT id, dominion_id, prestige, peasants, morale, updated_at
FROM dominion_ticks
WHERE dominion_id = $1
ORDER BY updated_at DESC NULLS LAST, id DESC
LIMIT 1;
