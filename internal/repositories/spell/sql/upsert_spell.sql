WITH upsert_spell AS (
  INSERT INTO spells
    (key, name, category, cost_mana, cost_strength, duration, cooldown, active, races)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  ON CONFLICT (key) DO UPDATE SET
    name=$2, category=$3, cost_mana=$4, cost_strength=$5,
    duration=$6, cooldown=$7, active=$8, races=$9,
    updated_at=now()
  RETURNING id
),
input_perks AS (
  SELECT us.id AS spell_id, p.key, p.value
  FROM upsert_spell us
  -- jsonb_to_recordset returns record; provide a column definition list so Postgres
  -- knows the names and types of the returned columns.
  CROSS JOIN LATERAL jsonb_to_recordset(COALESCE($10::jsonb, '[]'::jsonb)) AS p(key text, value text)
),
upsert_types AS (
  -- Insert any new spell perk types; do nothing on conflict
  INSERT INTO spell_perk_types (key)
  SELECT DISTINCT ip.key FROM input_perks ip
  ON CONFLICT (key) DO NOTHING
  RETURNING id, key
),
all_types AS (
  -- Combine newly inserted types with existing types for the input keys
  SELECT id, key FROM upsert_types
  UNION
  SELECT spt.id, spt.key
  FROM spell_perk_types spt
  WHERE spt.key IN (SELECT DISTINCT key FROM input_perks)
),
upsert_perks AS (
  INSERT INTO spell_perks (spell_id, spell_perk_type_id, value)
  SELECT ip.spell_id, at.id, ip.value
  FROM input_perks ip
  JOIN all_types at ON at.key = ip.key
  ON CONFLICT (spell_id, spell_perk_type_id) DO UPDATE SET value = EXCLUDED.value
  RETURNING 1
)
SELECT id FROM upsert_spell;
