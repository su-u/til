show variables like 'slow_query%';

show variables like 'long%';

set global slow_query_log=1;
set global slow_query_log_file='/tmp/slow.log';
set global long_query_time=1;


SELECT SLEEP(3);

SELECT SLEEP(10);

SELECT SLEEP(0.5);

SELECT * FROM address
WHERE address2  IS NOT NULL;

SHOW GLOBAL VARIABLES;

ALTER TABLE actor DROP COLUMN c1;


BEGIN;
SELECT COUNT(*) FROM address;

ALTER TABLE actor ADD COLUMN c1 INT;

COMMIT;

SELECT IS_FREE_LOCK('str');
SELECT GET_LOCK('str', 5);
SELECT RELEASE_LOCK('str');

SELECT trx_rows_locked FROM information_schema.INNODB_TRX;

SELECT * FROM sys.innodb_lock_waits;
SELECT * FROM performance_schema.data_locks;

SHOW ENGINE INNODB STATUS