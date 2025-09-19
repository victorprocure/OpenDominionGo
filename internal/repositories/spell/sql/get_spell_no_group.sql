SELECT
  s.id, s.key, s.name, s.category, s.cost_mana, s.cost_strength,
  s.duration, s.cooldown, s.active, s.created_at, s.updated_at,
            COALESCE(
                (
                 SELECT json_agg(json_build_object(
                 		'id', r.id, 
                 		'key', r.key,
                        'name', r.name,
                        'alignment', r.alignment,
                        'home_land_type', r.home_land_type,
                        'description', r.description,
                        'playable', r.playable,
                        'attacker_difficulty', r.attacker_difficulty,
                        'explorer_difficulty', r.explorer_difficulty,
                        'converter_difficulty', r.converter_difficulty,
                        'overall_difficulty', r.overall_difficulty,
                        'created_at', r.created_at, 
                        'updated_at', r.updated_at))
                 FROM races r
                 %s
                ), '[]'::json
            ) AS races_json,
            COALESCE(
                json_agg(
                  json_build_object(
                    'id', sp.id, 
                    'value', sp.value,
                    'created_at', sp.created_at,
                    'updated_at', sp.updated_at,
                    'perk_type', json_build_object(
                                    'id', spt.id,
                                    'key', spt.key,
                                    'created_at', spt.created_at,
                                    'updated_at', spt.updated_at)
                  )
                ) FILTER (WHERE sp.id IS NOT NULL),
                '[]'::json
            ) AS perks_json
        FROM spells s
        LEFT JOIN spell_perks sp ON sp.spell_id = s.id
        LEFT JOIN spell_perk_types spt ON spt.id = sp.spell_perk_type_id
        %s
        GROUP BY s.id
