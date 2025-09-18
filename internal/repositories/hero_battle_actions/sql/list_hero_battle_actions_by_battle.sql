SELECT id, hero_battle_id, combatant_id, target_combatant_id, turn, action, damage, health, description
FROM hero_battle_actions
WHERE hero_battle_id = $1
ORDER BY turn ASC, id ASC;
