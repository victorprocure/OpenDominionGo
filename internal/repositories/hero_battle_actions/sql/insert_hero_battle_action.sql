INSERT INTO hero_battle_actions (hero_battle_id, combatant_id, target_combatant_id, turn, action, damage, health, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;
