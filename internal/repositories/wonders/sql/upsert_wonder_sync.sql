WITH upsert_wonder AS (
  INSERT INTO wonders (key, name, power, active)
  VALUES ($1, $2, $3, $4)
  ON CONFLICT (key) DO UPDATE
    SET name = EXCLUDED.name,
        power = EXCLUDED.power,
        active = EXCLUDED.active,
        updated_at = now()
  RETURNING id
),
input_perks AS (
  SELECT uw.id AS wonder_id, p.key, p.value
  FROM upsert_wonder uw
  CROSS JOIN LATERAL jsonb_to_recordset(COALESCE($5::jsonb, '[]'::jsonb)) AS p(key text, value text)
),
upsert_types AS (
  INSERT INTO wonder_perk_types (key)
  SELECT DISTINCT key FROM input_perks
  ON CONFLICT (key) DO UPDATE SET key = EXCLUDED.key
  RETURNING id, key
),
all_types AS (
  SELECT id, key FROM upsert_types
  UNION ALL
  SELECT wpt.id, wpt.key
  FROM wonder_perk_types wpt
  JOIN input_perks ip ON ip.key = wpt.key
),
upsert_perks AS (
  INSERT INTO wonder_perks (wonder_id, wonder_perk_type_id, value)
  SELECT ip.wonder_id, at.id, ip.value
  FROM input_perks ip
  JOIN all_types at ON at.key = ip.key
  ON CONFLICT (wonder_id, wonder_perk_type_id) DO UPDATE
    SET value = EXCLUDED.value
  RETURNING 1
)
SELECT id FROM upsert_wonder;