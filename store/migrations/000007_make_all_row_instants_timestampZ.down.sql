DO $$
DECLARE
  r  record;
  tz text := 'UTC'; -- the zone your values should be interpreted in
BEGIN
  FOR r IN
    SELECT c.table_schema, c.table_name, c.column_name
    FROM information_schema.columns c
    JOIN information_schema.tables t
      ON t.table_schema = c.table_schema AND t.table_name = c.table_name
    WHERE t.table_type = 'BASE TABLE'
      AND c.table_schema NOT IN ('pg_catalog','information_schema')
      -- AND c.table_schema = 'public' -- uncomment to limit to public
      AND c.column_name IN ('created_at','updated_at','deleted_at')
      AND c.data_type = 'timestamp with time zone'
  LOOP
    EXECUTE format(
      'ALTER TABLE %I.%I ALTER COLUMN %I TYPE timestamp without time zone USING %I AT TIME ZONE %L',
      r.table_schema, r.table_name, r.column_name, r.column_name, tz
    );
  END LOOP;
END$$;