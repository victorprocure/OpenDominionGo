SELECT id, round_league_id, number, name, start_date, end_date,
       created_at, updated_at, realm_size, pack_size, players_per_race,
       mixed_alignment, offensive_actions_prohibited_at, discord_guild_id,
       tech_version, largest_hit, assignment_complete
FROM rounds
WHERE start_date <= $1 AND end_date >= $1
ORDER BY start_date DESC, id DESC;