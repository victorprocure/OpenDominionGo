-- Upsert a tech and its perks (and perk types) in a single statement
-- Parameters:
--   $1  = key (text)
--   $2  = name (text)
--   $3  = prerequisites (text) -- comma-delimited or your chosen format
--   $4  = active (boolean)
--   $5  = version (integer)
--   $6  = x (integer)
--   $7  = y (integer)
--   $8  = perks JSON (jsonb) - object mapping of perk_key -> value, e.g. {"atk":"5","def":"2"}
--         Pass NULL or '{}' when no perks are provided

WITH upsert_tech AS (
  INSERT INTO techs (
    key, name, prerequisites, active, version, x, y
  )
  VALUES ($1, $2, $3, $4, $5, $6, $7)
  ON CONFLICT (key) DO UPDATE
    SET name          = EXCLUDED.name,
        prerequisites = EXCLUDED.prerequisites,
        active        = EXCLUDED.active,
        version       = EXCLUDED.version,
        x             = EXCLUDED.x,
        y             = EXCLUDED.y,
        updated_at    = now()
  RETURNING id
),
input_perks AS (
  -- Flatten the input perks JSON object for this tech into rows
  SELECT ut.id AS tech_id, p.key AS perk_key, p.value
  FROM upsert_tech ut
  CROSS JOIN LATERAL jsonb_each_text(COALESCE($8::jsonb, '{}'::jsonb)) AS p(key, value)
),
upsert_types AS (
  -- Ensure all perk types exist. Insert only new keys; do nothing on conflict.
  INSERT INTO tech_perk_types (key)
  SELECT DISTINCT ip.perk_key
  FROM input_perks ip
  ON CONFLICT (key) DO NOTHING
  RETURNING id, key
),
all_types AS (
  -- Resolve type IDs whether newly inserted or pre-existing.
  -- We take the ids returned from the insert and union them with the existing
  -- rows for the distinct perk keys in the input. Using UNION here prevents
  -- duplicate rows for the same key.
  SELECT id, key FROM upsert_types
  UNION
  SELECT tpt.id, tpt.key
  FROM tech_perk_types tpt
  WHERE tpt.key IN (SELECT DISTINCT ip.perk_key FROM input_perks ip)
),
upsert_perks AS (
  -- Upsert perk values keyed by (tech_id, tech_perk_type_id)
  INSERT INTO tech_perks (tech_id, tech_perk_type_id, value)
  SELECT ip.tech_id, at.id, ip.value
  FROM input_perks ip
  JOIN all_types at ON at.key = ip.perk_key
  ON CONFLICT (tech_id, tech_perk_type_id) DO UPDATE
    SET value      = EXCLUDED.value,
        updated_at = now()
  RETURNING 1
)
SELECT id FROM upsert_tech;
