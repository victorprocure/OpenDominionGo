WITH upsert AS (
  INSERT INTO daily_rankings (
    round_id, dominion_id, dominion_name, race_name, realm_number, realm_name,
    key, value, rank, previous_rank
  )
  VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
  ON CONFLICT (dominion_id, key) DO UPDATE
    SET round_id = EXCLUDED.round_id,
        dominion_name = EXCLUDED.dominion_name,
        race_name = EXCLUDED.race_name,
        realm_number = EXCLUDED.realm_number,
        realm_name = EXCLUDED.realm_name,
        value = EXCLUDED.value,
        rank = EXCLUDED.rank,
        previous_rank = EXCLUDED.previous_rank,
        updated_at = now()
  RETURNING id
)
SELECT id FROM upsert;
