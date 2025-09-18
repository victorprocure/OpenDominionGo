DO $$
DECLARE
  r record;
BEGIN
  FOR r IN
    SELECT c.table_schema, c.table_name
    FROM information_schema.columns c
    JOIN information_schema.tables t
      ON t.table_schema = c.table_schema AND t.table_name = c.table_name
    WHERE t.table_type = 'BASE TABLE'
      AND c.column_name = 'key'
      AND c.table_schema NOT IN ('pg_catalog','information_schema')
      AND EXISTS (
        SELECT 1 FROM information_schema.columns x
        WHERE x.table_schema = c.table_schema
          AND x.table_name   = c.table_name
          AND x.column_name  = 'created_at'
          AND x.data_type IN ('timestamp without time zone','timestamp with time zone')
      )
      AND EXISTS (
        SELECT 1 FROM information_schema.columns y
        WHERE y.table_schema = c.table_schema
          AND y.table_name   = c.table_name
          AND y.column_name  = 'updated_at'
          AND y.data_type IN ('timestamp without time zone','timestamp with time zone')
      )
  LOOP
    -- Drop the UNIQUE constraint added by the up migration (if present)
    EXECUTE format('ALTER TABLE %I.%I DROP CONSTRAINT IF EXISTS %I',
                   r.table_schema, r.table_name, r.table_name || '_key_unique');

    -- Revert NOT NULL on key (best-effort; original state is unknown)
    EXECUTE format('ALTER TABLE %I.%I ALTER COLUMN key DROP NOT NULL',
                   r.table_schema, r.table_name);
  END LOOP;
END$$;