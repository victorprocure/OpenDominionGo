WITH upsert_race AS (
  INSERT INTO races (
    key, name, alignment, description,
    attacker_difficulty, explorer_difficulty, converter_difficulty, overall_difficulty,
    home_land_type, playable
  )
  VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
  ON CONFLICT (key) DO UPDATE SET
    name = EXCLUDED.name,
    alignment = EXCLUDED.alignment,
    description = EXCLUDED.description,
    attacker_difficulty = EXCLUDED.attacker_difficulty,
    explorer_difficulty = EXCLUDED.explorer_difficulty,
    converter_difficulty = EXCLUDED.converter_difficulty,
    overall_difficulty = EXCLUDED.overall_difficulty,
    home_land_type = EXCLUDED.home_land_type,
    playable = EXCLUDED.playable,
    updated_at = now()
  RETURNING id
),
race_perks_input AS (
  SELECT ur.id AS race_id, p.key, p.value
  FROM upsert_race ur
  CROSS JOIN LATERAL jsonb_to_recordset(COALESCE($11::jsonb, '[]'::jsonb)) AS p(key text, value text)
),
upsert_race_perk_types AS (
  INSERT INTO race_perk_types (key)
  SELECT DISTINCT key FROM race_perks_input
  ON CONFLICT (key) DO UPDATE SET key = EXCLUDED.key
  RETURNING id, key
),
all_race_perk_types AS (
  SELECT id, key FROM upsert_race_perk_types
  UNION ALL
  SELECT rpt.id, rpt.key
  FROM race_perk_types rpt
  JOIN race_perks_input i ON i.key = rpt.key
),
upsert_race_perks AS (
  INSERT INTO race_perks (race_id, race_perk_type_id, value)
  SELECT i.race_id, t.id, i.value
  FROM race_perks_input i
  JOIN all_race_perk_types t USING (key)
  ON CONFLICT (race_id, race_perk_type_id) DO UPDATE
    SET value = EXCLUDED.value
  RETURNING 1
),
units_input AS (
  SELECT
    ur.id AS race_id,
    u.name,
    u.type,
    COALESCE(u.need_boat, false) AS need_boat,
    COALESCE(u.cost_platinum, 0) AS cost_platinum,
    COALESCE(u.cost_ore, 0) AS cost_ore,
    COALESCE(u.cost_lumber, 0) AS cost_lumber,
    COALESCE(u.cost_gems, 0) AS cost_gems,
    COALESCE(u.cost_mana, 0) AS cost_mana,
    COALESCE(u.power_offense, 0) AS power_offense,
    COALESCE(u.power_defense, 0) AS power_defense,
    u.perks AS perks_json
  FROM upsert_race ur
  CROSS JOIN LATERAL jsonb_to_recordset(COALESCE($12::jsonb, '[]'::jsonb)) AS u(
    name text,
    type text,
    need_boat boolean,
    cost_platinum int,
    cost_ore int,
    cost_lumber int,
    cost_gems int,
    cost_mana int,
    power_offense int,
    power_defense int,
    perks jsonb
  )
),
upsert_units AS (
  INSERT INTO units (
    race_id, name, unit_type, need_boat,
    cost_platinum, cost_ore, cost_lumber, cost_gems, cost_mana,
    power_offense, power_defense
  )
  SELECT
    race_id, name, u.type, need_boat,
    cost_platinum, cost_ore, cost_lumber, cost_gems, cost_mana,
    power_offense, power_defense
  FROM units_input u
  ON CONFLICT (race_id, name) DO UPDATE SET
    unit_type = EXCLUDED.unit_type,
    need_boat = EXCLUDED.need_boat,
    cost_platinum = EXCLUDED.cost_platinum,
    cost_ore = EXCLUDED.cost_ore,
    cost_lumber = EXCLUDED.cost_lumber,
    cost_gems = EXCLUDED.cost_gems,
    cost_mana = EXCLUDED.cost_mana,
    power_offense = EXCLUDED.power_offense,
    power_defense = EXCLUDED.power_defense,
    updated_at = now()
  RETURNING id, race_id, name
),
unit_perks_input AS (
  SELECT uu.id AS unit_id, p.key, p.value
  FROM upsert_units uu
  JOIN units_input ui ON ui.race_id = uu.race_id AND ui.name = uu.name
  CROSS JOIN LATERAL jsonb_to_recordset(COALESCE(ui.perks_json, '[]'::jsonb)) AS p(key text, value text)
),
upsert_unit_perk_types AS (
  INSERT INTO unit_perk_types (key)
  SELECT DISTINCT key FROM unit_perks_input
  ON CONFLICT (key) DO UPDATE SET key = EXCLUDED.key
  RETURNING id, key
),
all_unit_perk_types AS (
  SELECT id, key FROM upsert_unit_perk_types
  UNION ALL
  SELECT upt.id, upt.key
  FROM unit_perk_types upt
  JOIN unit_perks_input i ON i.key = upt.key
),
upsert_unit_perks AS (
  INSERT INTO unit_perks (unit_id, unit_perk_type_id, value)
  SELECT i.unit_id, t.id, i.value
  FROM unit_perks_input i
  JOIN all_unit_perk_types t USING (key)
  ON CONFLICT (unit_id, unit_perk_type_id) DO UPDATE
    SET value = EXCLUDED.value
  RETURNING 1
)
SELECT id FROM upsert_race;