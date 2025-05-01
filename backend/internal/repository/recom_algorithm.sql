-- Function to generate action points based on params
-- Already Stored in CockroachDB
CREATE OR REPLACE FUNCTION gen_rating(rating_to VARCHAR(25), target_from DECIMAL(5,2), target_to DECIMAL(5,2), action VARCHAR(25))
RETURNS INTEGER AS $$
DECLARE 
points INTEGER:=0; -- acum of points: higher is better
diffPerc NUMERIC;
BEGIN
  IF rating_to ILIKE 'str%' THEN --Strong-Buy
    points:=points+4;
  ELSIF rating_to ILIKE 'buy' THEN -- Buy
    points:=points+2;
  END IF;

  IF target_to=target_from THEN -- If target keeps, fine signal
    points:=points+1;
  ELSE -- there is a variation 
    diffPerc:=ROUND(( (target_to-target_from)*100 )/target_from,2); -- diff %
    IF diffPerc>0 THEN -- diff +% (good)
      IF diffPerc<=5 THEN
        points:=points+1;
      ELSIF diffPerc<=15 THEN
        points:=points+2;
      ELSIF diffPerc<=30 THEN
        points:=points+3;
      ELSIF diffPerc<=50 THEN
        points:=points+4;
      ELSE
        points:=points+5;
      END IF;
    ELSE -- diff -% (bad)
      IF diffPerc>=-5 THEN
        points:=points-1;
      ELSIF diffPerc>=-15 THEN
        points:=points-2;
      ELSIF diffPerc>=-30 THEN
        points:=points-3;
      ELSIF diffPerc>=-50 THEN
        points:=points-4;
      ELSE
        points:=points-5;
      END IF;
    END IF;
  END IF;

  -- Quantify the 'action' impact
  IF action ILIKE 'reiterated%' OR action ILIKE 'initiated%' OR action ILIKE '%set%' THEN
    points:=points+1;
  ELSIF action ILIKE '%raised%' OR action ILIKE '%upgraded%' THEN
    points:=points+2;
  ELSIF action ILIKE '%lowered%' OR action ILIKE '%downgraded%' THEN
    points:=points-2;
  END IF;

  RETURN points -- Acum results
END;
$$ LANGUAGE PLpgSQL;


-- Base Query to get the best actions to invest (invoke function above)
-- NOTE: 'WHERE total_rating>5' is adaptable; Highest comparition for more specificity, lower for less
SELECT * FROM (SELECT gen_rating(rating_to, target_from, target_to, action) AS total_rating,target_from, target_to, action, rating_to FROM STOCK WHERE (rating_to ILIKE '%buy%') AND (rating_to NOT ILIKE '%spe%') AND (target_from>0 AND target_to>0))as sub WHERE total_rating>5 ORDER BY total_rating DESC;