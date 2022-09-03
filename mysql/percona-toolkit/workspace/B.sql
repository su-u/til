SELECT *
FROM actor
WHERE actor_id = 1;


SELECT *
FROM actor
WHERE actor_id = 1
FOR UPDATE ;

BEGIN;

SELECT * FROM actor WHERE actor_id = 1 LOCK IN SHARE MODE;

COMMIT;
