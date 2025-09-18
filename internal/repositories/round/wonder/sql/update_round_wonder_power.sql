UPDATE round_wonders
SET power = $2,
    updated_at = now()
WHERE id = $1;
