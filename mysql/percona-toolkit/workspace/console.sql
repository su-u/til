show variables like 'slow_query%';

show variables like 'long%';

set global slow_query_log_file='/tmp/slow.log';
set global slow_query_log=1;
set global long_query_time=1;


SELECT SLEEP(3);

SELECT * FROM address
WHERE address2  IS NOT NULL