INSERT INTO telescope_monitoring (tag)
VALUES ($1)
ON CONFLICT DO NOTHING;
