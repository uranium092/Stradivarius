-- Function to generate action points based on params
-- Already Stored in CockroachDB
CREATE OR REPLACE FUNCTION gen_rating(rating_to VARCHAR(25), target_from DECIMAL(5,2), target_to DECIMAL(5,2), action VARCHAR(25))
RETURNS NUMERIC AS $$
    SELECT 
    (CASE
      WHEN rating_to='Strong-Buy' THEN 4
      ELSE 2
    END * 0.35) +
    (CASE
      WHEN target_to=target_from THEN 0.1
      WHEN ROUND(((target_to - target_from) * 100) / target_from)>0 THEN ROUND(POWER((((target_to - target_from) * 100.0) / target_from) / 10.0, 1.2))
      ELSE
        -ROUND(POWER(ABS(((target_to - target_from) * 100.0) / target_from) / 10.0, 1.2))
    END * 0.45) +
    (CASE
      WHEN action ILIKE 'reiterated%' OR action ILIKE 'initiated%' OR action ILIKE '%set%' THEN 1
      WHEN action ILIKE '%raised%' OR action ILIKE '%upgraded%' THEN 2
      ELSE -2
    END * 0.20)
$$ LANGUAGE sql;

-- Base Query to get the best actions to invest (invoke function above)
-- NOTE: 'WHERE total_rating>1.5' is adaptable; Highest comparition for more specificity, lower for less
SELECT * FROM (SELECT gen_rating(rating_to, target_from, target_to, action) AS total_rating,* FROM STOCK WHERE (rating_to='Strong-Buy' OR rating_to='Buy') AND (target_from>0 AND target_to>0))as sub WHERE total_rating>1.5 ORDER BY total_rating DESC, id;