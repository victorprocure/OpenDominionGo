-- Drop triggers that auto-update updated_at on all base tables
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT c.table_schema, c.table_name
    FROM information_schema.columns c
    JOIN information_schema.tables t
      ON t.table_schema = c.table_schema AND t.table_name = c.table_name
    WHERE c.column_name = 'updated_at'
      AND t.table_type = 'BASE TABLE'
      AND c.table_schema NOT IN ('pg_catalog','information_schema')
      AND c.data_type IN ('timestamp without time zone','timestamp with time zone')
  LOOP
    EXECUTE format('DROP TRIGGER IF EXISTS %I ON %I.%I',
                   'trg_'||r.table_name||'_set_updated_at', r.table_schema, r.table_name);
  END LOOP;
END$$;

-- Drop triggers that protect created_at on all base tables
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT c.table_schema, c.table_name
    FROM information_schema.columns c
    JOIN information_schema.tables t
      ON t.table_schema = c.table_schema AND t.table_name = c.table_name
    WHERE c.column_name = 'created_at'
      AND t.table_type = 'BASE TABLE'
      AND c.table_schema NOT IN ('pg_catalog','information_schema')
      AND c.data_type IN ('timestamp without time zone','timestamp with time zone')
  LOOP
    EXECUTE format('DROP TRIGGER IF EXISTS %I ON %I.%I',
                   'trg_'||r.table_name||'_protect_created_at', r.table_schema, r.table_name);
  END LOOP;
END$$;

-- Remove DEFAULT now() from created_at on all base tables (best-effort; original defaults are not restored)
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT c.table_schema, c.table_name
    FROM information_schema.columns c
    JOIN information_schema.tables t
      ON t.table_schema = c.table_schema AND t.table_name = c.table_name
    WHERE c.column_name = 'created_at'
      AND t.table_type = 'BASE TABLE'
      AND c.table_schema NOT IN ('pg_catalog','information_schema')
      AND c.data_type IN ('timestamp without time zone','timestamp with time zone')
  LOOP
    EXECUTE format('ALTER TABLE %I.%I ALTER COLUMN created_at DROP DEFAULT', r.table_schema, r.table_name);
  END LOOP;
END$$;

-- Remove DEFAULT now() from updated_at on all base tables
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT c.table_schema, c.table_name
    FROM information_schema.columns c
    JOIN information_schema.tables t
      ON t.table_schema = c.table_schema AND t.table_name = c.table_name
    WHERE c.column_name = 'updated_at'
      AND t.table_type = 'BASE TABLE'
      AND c.table_schema NOT IN ('pg_catalog','information_schema')
      AND c.data_type IN ('timestamp without time zone','timestamp with time zone')
  LOOP
    EXECUTE format('ALTER TABLE %I.%I ALTER COLUMN updated_at DROP DEFAULT', r.table_schema, r.table_name);
  END LOOP;
END$$;

-- Drop trigger functions (safe if not used elsewhere)
DROP FUNCTION IF EXISTS set_updated_at();
DROP FUNCTION IF EXISTS protect_created_at();