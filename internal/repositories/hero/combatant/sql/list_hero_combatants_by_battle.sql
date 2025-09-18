SELECT id, hero_battle_id, hero_id, dominion_id, name, health, attack, defense, evasion, focus, counter, recover, current_health, level
FROM hero_combatants
WHERE hero_battle_id = $1
ORDER BY id ASC;
