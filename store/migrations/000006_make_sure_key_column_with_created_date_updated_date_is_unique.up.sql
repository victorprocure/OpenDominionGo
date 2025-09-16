DO $$
DECLARE
  r record;
  has_dups bool;
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
    -- Ensure key is NOT NULL
    EXECUTE format('ALTER TABLE %I.%I ALTER COLUMN key SET NOT NULL', r.table_schema, r.table_name);

    -- Skip if a UNIQUE constraint on (key) already exists
    PERFORM 1
    FROM pg_constraint con
    JOIN pg_class rel ON rel.oid = con.conrelid
    JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
    WHERE con.contype = 'u'
      AND nsp.nspname = r.table_schema
      AND rel.relname = r.table_name
      AND con.conkey IS NOT NULL
      AND array_length(con.conkey,1) = 1
      AND (SELECT attname FROM pg_attribute WHERE attrelid = rel.oid AND attnum = con.conkey[1]) = 'key';

    IF NOT FOUND THEN
      -- Fail early if duplicates exist
      EXECUTE format('SELECT EXISTS (SELECT 1 FROM %I.%I GROUP BY key HAVING COUNT(*) > 1)', r.table_schema, r.table_name)
        INTO has_dups;
      IF has_dups THEN
        RAISE EXCEPTION 'Cannot add UNIQUE(key) on %.%: duplicate keys present', r.table_schema, r.table_name;
      END IF;

      -- Add UNIQUE constraint
      EXECUTE format('ALTER TABLE %I.%I ADD CONSTRAINT %I UNIQUE (key)',
                     r.table_schema, r.table_name, r.table_name || '_key_unique');
    END IF;
  END LOOP;
END$$;