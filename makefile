include .env

create_container:
	@docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${DB_DOCKER_USER} -e POSTGRES_PASSWORD=${DB_DOCKER_PASSWORD} -d postgres:16-alpine

start_container:
	@docker start ${DB_DOCKER_CONTAINER}

create_db:
	@docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${DB_DOCKER_USER} --owner=${DB_DOCKER_USER} ${DB_NAME}

drop_db:
	@docker exec -it ${DB_DOCKER_CONTAINER} dropdb ${DB_NAME}

open_db:
	@docker exec -it ${DB_DOCKER_CONTAINER} psql -U ${DB_DOCKER_USER}

migrateup:
	migrate -path database/migration -database "postgresql://${DB_DOCKER_USER}:${DB_DOCKER_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://${DB_DOCKER_USER}:${DB_DOCKER_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down

migrateforce:
	migrate -path database/migration -database "postgresql://${DB_DOCKER_USER}:${DB_DOCKER_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" force 1

run:
	@go build -o api main.go && ./api