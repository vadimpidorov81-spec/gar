include .env
export

.PHONY: migrate-up migrate-down migrate-version migrate-force migrate-create

run:
	DB_URL=${DB_URL} go run main.go

clean gar_addr_obj:

migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_URL} up

migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_URL} down 1

migrate-version:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_URL} version

