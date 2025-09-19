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
  -- Expecting a JSON object mapping of perk_key -> value, e.g. {"mana_production": "5"}
  SELECT ur.id AS race_id, p.key, p.value
  FROM upsert_race ur
  CROSS JOIN LATERAL jsonb_each_text(COALESCE($11::jsonb, '{}'::jsonb)) AS p
),
upsert_race_perk_types AS (
  -- Insert any new race perk types; do nothing on conflict to avoid updating the same key twice
  INSERT INTO race_perk_types (key)
  SELECT DISTINCT key FROM race_perks_input
  ON CONFLICT (key) DO NOTHING
  RETURNING id, key
),
all_race_perk_types AS (
  -- Combine newly inserted types with existing types for the input keys
  SELECT id, key FROM upsert_race_perk_types
  UNION
  SELECT rpt.id, rpt.key
  FROM race_perk_types rpt
  WHERE rpt.key IN (SELECT DISTINCT key FROM race_perks_input)
),
upsert_race_perks AS (
  INSERT INTO race_perks (race_id, race_perk_type_id, value)
  SELECT i.race_id, t.id, (i.value)::double precision
  FROM race_perks_input i
  JOIN all_race_perk_types t USING (key)
  ON CONFLICT (race_id, race_perk_type_id) DO UPDATE
    SET value = EXCLUDED.value
  RETURNING 1
),
units_input AS (
  -- Unnest units array and extract fields; cast values to the expected column types
  SELECT
    ur.id AS race_id,
    (elem ->> 'name') AS name,
    (elem ->> 'type') AS type,
    ((ord)::text) AS slot,
    COALESCE((elem ->> 'need_boat')::boolean, false) AS need_boat,
    COALESCE((elem ->> 'cost_platinum')::int, 0) AS cost_platinum,
    COALESCE((elem ->> 'cost_ore')::int, 0) AS cost_ore,
    COALESCE((elem ->> 'cost_lumber')::int, 0) AS cost_lumber,
    COALESCE((elem ->> 'cost_gems')::int, 0) AS cost_gems,
    COALESCE((elem ->> 'cost_mana')::int, 0) AS cost_mana,
    COALESCE((elem ->> 'power_offense')::double precision, 0) AS power_offense,
    COALESCE((elem ->> 'power_defense')::double precision, 0) AS power_defense,
    elem -> 'perks' AS perks_json
  FROM upsert_race ur
  CROSS JOIN LATERAL jsonb_array_elements(COALESCE($12::jsonb, '[]'::jsonb)) WITH ORDINALITY AS elem(elem, ord)
),
upsert_units AS (
  INSERT INTO units (
    race_id, name, slot, type, need_boat,
    cost_platinum, cost_ore, cost_lumber, cost_gems, cost_mana,
    power_offense, power_defense
  )
  SELECT
    race_id, name, slot, u.type, need_boat,
    cost_platinum, cost_ore, cost_lumber, cost_gems, cost_mana,
    power_offense, power_defense
  FROM units_input u
  ON CONFLICT (race_id, slot) DO UPDATE SET
    type = EXCLUDED.type,
    need_boat = EXCLUDED.need_boat,
    cost_platinum = EXCLUDED.cost_platinum,
    cost_ore = EXCLUDED.cost_ore,
    cost_lumber = EXCLUDED.cost_lumber,
    cost_gems = EXCLUDED.cost_gems,
    cost_mana = EXCLUDED.cost_mana,
    power_offense = EXCLUDED.power_offense,
    power_defense = EXCLUDED.power_defense,
    updated_at = now()
  RETURNING id, race_id, name, slot
),
unit_perks_input AS (
  SELECT uu.id AS unit_id, p.key, p.value
  FROM upsert_units uu
  JOIN units_input ui ON ui.race_id = uu.race_id AND ui.slot = uu.slot
  -- Only call jsonb_each_text when perks_json is actually an object
  CROSS JOIN LATERAL (
    SELECT key, value
    FROM jsonb_each_text(ui.perks_json)
    WHERE jsonb_typeof(ui.perks_json) = 'object'
  ) AS p
),
upsert_unit_perk_types AS (
  -- Insert any new unit perk types; do nothing on conflict
  INSERT INTO unit_perk_types (key)
  SELECT DISTINCT key FROM unit_perks_input
  ON CONFLICT (key) DO NOTHING
  RETURNING id, key
),
all_unit_perk_types AS (
  -- Combine newly inserted types with existing types for the input keys
  SELECT id, key FROM upsert_unit_perk_types
  UNION
  SELECT upt.id, upt.key
  FROM unit_perk_types upt
  WHERE upt.key IN (SELECT DISTINCT key FROM unit_perks_input)
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
