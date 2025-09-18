-- Upsert a hero upgrade and its perks in one go
-- $1 key, $2 name, $3 level, $4 type, $5 icon, $6 classes (text), $7 active (bool)
-- $8 perks jsonb object mapping (key->value)
-- NOTE: Perks present in DB but NOT in $8 will be deleted to match YAML exactly.
WITH upsert_upgrade AS (
  INSERT INTO hero_upgrades (key, name, level, type, icon, classes, active)
  VALUES ($1, $2, $3, $4, $5, $6, $7)
  ON CONFLICT (key) DO UPDATE
    SET name = EXCLUDED.name,
        level = EXCLUDED.level,
        type = EXCLUDED.type,
        icon = EXCLUDED.icon,
        classes = EXCLUDED.classes,
        active = EXCLUDED.active,
        updated_at = now()
  RETURNING id
),
input_perks AS (
  SELECT hu.id AS hero_upgrade_id, p.key, p.value
  FROM upsert_upgrade hu
  CROSS JOIN LATERAL jsonb_each_text(COALESCE($8::jsonb, '{}'::jsonb)) AS p(key, value)
),
upsert_perks AS (
  INSERT INTO hero_upgrade_perks (hero_upgrade_id, key, value)
  SELECT hero_upgrade_id, key, value FROM input_perks
  ON CONFLICT (hero_upgrade_id, key) DO UPDATE
    SET value = EXCLUDED.value,
        updated_at = now()
  RETURNING 1
),
delete_perks AS (
  -- Remove any existing perks for this hero upgrade that are not present in the incoming set
  DELETE FROM hero_upgrade_perks hup
  USING upsert_upgrade uu
  WHERE hup.hero_upgrade_id = uu.id
    AND NOT EXISTS (
      SELECT 1 FROM input_perks ip
      WHERE ip.hero_upgrade_id = uu.id AND ip.key = hup.key
    )
  RETURNING 1
)
SELECT id FROM upsert_upgrade;
