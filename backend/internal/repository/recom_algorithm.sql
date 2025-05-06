-- Function to generate action points based on params
-- Already Stored in CockroachDB
CREATE OR REPLACE FUNCTION gen_rating(rating_to VARCHAR(25), target_from DECIMAL(5,2), target_to DECIMAL(5,2), action VARCHAR(25))
RETURNS NUMERIC AS $$
DECLARE
  rating_point INTEGER := 0;
  equal_point INTEGER := 0;
  diffPerc NUMERIC := 0;
  variation_point NUMERIC := 0;
  action_point INTEGER := 0;
  final_points NUMERIC := 0;
BEGIN
  IF rating_to ILIKE 'str%' THEN --Strong-Buy
    rating_point:=4;
  ELSIF rating_to ILIKE 'buy' THEN -- Buy
    rating_point:=2;
  END IF;

  IF target_to = target_from THEN -- If target keeps, fine signal
    equal_point:=1;
  ELSE --there is a variation 
    diffPerc := ROUND(((target_to - target_from) * 100) / target_from); -- diff %
    IF diffPerc>0 THEN -- diff +% (good)
      variation_point:=ROUND(POWER((diffPerc / 10), 1.2));
    ELSE -- diff -% (bad)
      variation_point:=-ROUND(POWER(ABS(diffPerc)/ 10, 1.2));
    END IF;
  END IF;

  -- Quantify the 'action' impact
  IF action ILIKE 'reiterated%' OR action ILIKE 'initiated%' OR action ILIKE '%set%' THEN
    action_point:=1;
  ELSIF action ILIKE '%raised%' OR action ILIKE '%upgraded%' THEN
    action_point:=2;
  ELSIF action ILIKE '%lowered%' OR action ILIKE '%downgraded%' THEN
    action_point:=-2;
  END IF;

  -- 'ponderar' values. Each one has an special %
  final_points:=(variation_point*0.4)+(rating_point*0.3)+(action_point*0.15)+(equal_point*0.15);

  RETURN ROUND(final_points,2);
  
END;
$$ LANGUAGE PLpgSQL;


-- Base Query to get the best actions to invest (invoke function above)
-- NOTE: 'WHERE total_rating>1.5' is adaptable; Highest comparition for more specificity, lower for less
SELECT * FROM (SELECT gen_rating(rating_to, target_from, target_to, action) AS total_rating,* FROM STOCK WHERE (rating_to ILIKE '%buy%') AND (rating_to NOT ILIKE '%spe%') AND (target_from>0 AND target_to>0))as sub WHERE total_rating>1.5 ORDER BY total_rating DESC, id;