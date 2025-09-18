DROP INDEX IF EXISTS idx_spells_races_gin;

ALTER TABLE spells ADD COLUMN races_str text;

UPDATE spells
SET races_str = CASE
  WHEN races IS NULL THEN NULL
  WHEN array_length(races, 1) IS NULL THEN ''
  ELSE array_to_string(races, ', ')
END;

ALTER TABLE spells DROP COLUMN races;
ALTER TABLE spells RENAME COLUMN races_str TO races;
