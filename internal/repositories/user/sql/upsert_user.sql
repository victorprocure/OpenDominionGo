INSERT INTO users (
    activated,
    avatar,
    display_name,
    email,
    last_online,
    message_board_last_read,
    password,
    rating
    settings,
	skin,
)
VALUES (
    $1,   -- activated
    $2,   -- avatar
    $3,   -- display_name
    $4,   -- email
    $5,   -- last_online
    $6,   -- message_board_last_read
    $7,   -- password
    $8,   -- rating
    $9,   -- settings
    $10   -- skin
)
ON CONFLICT (email) DO UPDATE
SET
    activated               = EXCLUDED.activated,
	avatar                  = EXCLUDED.avatar,
    display_name            = EXCLUDED.display_name,
    email                   = EXCLUDED.email,
	last_online             = EXCLUDED.last_online,
    message_board_last_read = EXCLUDED.message_board_last_read,
    password                = EXCLUDED.password,
    rating                  = EXCLUDED.rating
	settings                = EXCLUDED.settings,
	skin                    = EXCLUDED.skin,
RETURNING id
