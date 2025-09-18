  -- Auto-update updated_at
CREATE OR REPLACE FUNCTION set_updated_at() RETURNS trigger AS $$
BEGIN
  NEW.updated_at := NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Prevent changing created_at
CREATE OR REPLACE FUNCTION protect_created_at() RETURNS trigger AS $$
BEGIN
  IF NEW.created_at IS DISTINCT FROM OLD.created_at THEN
    RAISE EXCEPTION 'created_at is immutable';
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Set DEFAULT now() on all tables that have created_at
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
    EXECUTE format('ALTER TABLE %I.%I ALTER COLUMN created_at SET DEFAULT now()', r.table_schema, r.table_name);
  END LOOP;
END$$;

-- Set DEFAULT now() on all tables that have updated_at
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
    EXECUTE format('ALTER TABLE %I.%I ALTER COLUMN updated_at SET DEFAULT now()', r.table_schema, r.table_name);
  END LOOP;
END$$;

-- Attach set_updated_at to every table that has an updated_at column
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
    EXECUTE format('CREATE TRIGGER %I
                    BEFORE UPDATE ON %I.%I
                    FOR EACH ROW
                    EXECUTE FUNCTION set_updated_at()',
                   'trg_'||r.table_name||'_set_updated_at', r.table_schema, r.table_name);
  END LOOP;
END$$;

-- Attach protect_created_at to every table that has a created_at column
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
    EXECUTE format('CREATE TRIGGER %I
                    BEFORE UPDATE ON %I.%I
                    FOR EACH ROW
                    EXECUTE FUNCTION protect_created_at()',
                   'trg_'||r.table_name||'_protect_created_at', r.table_schema, r.table_name);
  END LOOP;
END$$;