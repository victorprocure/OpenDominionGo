DO $$
DECLARE
  r  record;
  tz text := 'UTC'; -- change if your existing values are in another zone
BEGIN
  FOR r IN
    SELECT c.table_schema, c.table_name, c.column_name
    FROM information_schema.columns c
    JOIN information_schema.tables t
      ON t.table_schema = c.table_schema AND t.table_name = c.table_name
    WHERE t.table_type = 'BASE TABLE'
      AND c.table_schema NOT IN ('pg_catalog','information_schema')
      -- uncomment to limit to public schema:
      -- AND c.table_schema = 'public'
      AND c.column_name IN ('created_at','updated_at','deleted_at')
      AND c.data_type = 'timestamp without time zone'
  LOOP
    EXECUTE format(
      'ALTER TABLE %I.%I ALTER COLUMN %I TYPE timestamptz USING %I AT TIME ZONE %L',
      r.table_schema, r.table_name, r.column_name, r.column_name, tz
    );
  END LOOP;
END$$;