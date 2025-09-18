DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint WHERE conname = 'units_race_id_name_unique'
  ) THEN
    ALTER TABLE units ADD CONSTRAINT units_race_id_name_unique UNIQUE (race_id, name);
  END IF;
END
$$;