SELECT id, race_id, name, slot, type, need_boat,
       cost_platinum, cost_ore, cost_lumber, cost_gems, cost_mana,
        power_offense, power_defense
FROM units
WHERE race_id = $1
ORDER BY slot::int ASC;
