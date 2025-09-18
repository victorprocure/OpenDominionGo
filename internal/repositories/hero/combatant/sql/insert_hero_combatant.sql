INSERT INTO hero_combatants (hero_battle_id, hero_id, dominion_id, name, health, attack, defense, evasion, focus, counter, recover, level, current_health)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $5)
RETURNING id;
