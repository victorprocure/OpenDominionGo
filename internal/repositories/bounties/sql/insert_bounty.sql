INSERT INTO bounties (
  round_id,
  source_realm_id,
  source_dominion_id,
  target_dominion_id,
  collected_by_dominion_id,
  type,
  reward
) VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING id;
