fmt:
	./script/fmt.sh

build:
	docker compose build

up:
	docker compose up

ps:
	docker compose ps

test:
	go test -v ./...

migration:
	set -a && source .env.prod && set +a&&\
	atlas schema apply \
	-u "mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@$${MYSQL_HOST}:$${MYSQL_PORT}/$${MYSQL_DATABASE}" \
	--to file://db/atlas/schema.hcl

schema_output:
	mkdir -p db/atlas &&\
	set -a && source .env.prod && set +a&&\
	atlas schema inspect -u "mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@$${MYSQL_HOST}:$${MYSQL_PORT}/$${MYSQL_DATABASE}" > db/atlas/schema.hcl

sql_output:
	mkdir -p db/atlas &&\
	set -a && source .env.prod && set +a&&\
	atlas schema inspect -u "mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@$${MYSQL_HOST}:$${MYSQL_PORT}/$${MYSQL_DATABASE}" --format "{{ sql . \" \" }}" > db/atlas/schema.sql

lint:
	./script/lint.sh

gen_db:
	./script/gen_db.sh

gen_proto:
	buf generate
	rm -r gen/yuovision-proto
	go mod tidy

prod:
	set -a && source .env.prod && set +a&&\
	go run cmd/server/main.go
