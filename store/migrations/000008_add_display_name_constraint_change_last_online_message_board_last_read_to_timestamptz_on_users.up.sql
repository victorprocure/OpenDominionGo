CREATE EXTENSION IF NOT EXISTS citext;

ALTER TABLE users 
  ALTER COLUMN last_online TYPE timestamptz USING last_online AT TIME ZONE 'UTC',
  ALTER COLUMN message_board_last_read TYPE timestamptz USING message_board_last_read AT TIME ZONE 'UTC',
  ALTER COLUMN display_name TYPE citext USING display_name::citext,
  ALTER COLUMN email TYPE citext USING email::citext,
  ADD CONSTRAINT users_display_name_unique UNIQUE (display_name);