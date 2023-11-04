pt-query-digest:
	docker-compose exec toolkit bash -c "pt-query-digest /tmp/slow.log"
up:
	docker-compose up -d
bash:
	docker-compose exec toolkit bash

DSN=h=mysql,A=utf8mb4,u=root,p=root,D=sakila,t=actor,P=3306

pt-change-dry:
	docker-compose exec toolkit bash -c \
		"pt-online-schema-change --alter \"ADD COLUMN c1 INT\" ${DSN} --dry-run --alter-foreign-keys-method auto --recursion-method processlist"

pt-change-execute:
	docker-compose exec toolkit bash -c \
		"pt-online-schema-change --alter \"ADD COLUMN c1 INT\" ${DSN} --execute --alter-foreign-keys-method auto --recursion-method processlist"