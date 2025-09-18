SELECT COALESCE(SUM(amount), 0)
FROM valor
WHERE dominion_id = $1;
