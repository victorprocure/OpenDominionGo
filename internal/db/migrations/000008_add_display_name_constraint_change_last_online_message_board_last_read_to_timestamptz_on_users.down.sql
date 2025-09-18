ALTER TABLE users 
  DROP CONSTRAINT IF EXISTS users_display_name_unique,
  ALTER COLUMN email type VARCHAR,
  ALTER COLUMN display_name TYPE VARCHAR(191),
  ALTER COLUMN last_online TYPE timestamp WITHOUT TIME ZONE USING last_online AT TIME ZONE 'UTC',
  ALTER COLUMN message_board_last_read TYPE timestamp WITHOUT TIME ZONE USING message_board_last_read AT TIME ZONE 'UTC';