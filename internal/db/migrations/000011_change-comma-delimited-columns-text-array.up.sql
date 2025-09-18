ALTER TABLE spells ADD COLUMN races_arr text[];

UPDATE spells
SET races_arr = CASE
  WHEN COALESCE(races, '') = '' THEN NULL
  ELSE regexp_split_to_array(races, '\s*,\s*')::text[]
END;

ALTER TABLE spells DROP COLUMN races;
ALTER TABLE spells RENAME COLUMN races_arr TO races;

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_spells_races_gin
  ON spells USING GIN (races);