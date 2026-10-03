include ./.env

DB_URL=postgres://$(DB_USER):$(DB_PASS):$

migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH)

migrate-down:
	@migrate -database $(DB_URL )

print-db-url:
@echo

seed-persons:
@psql $(DB_URL) < $(SEEDER_PATH)

seed-list :
@for r %i in ($(SEEDER_PATH)) do echo %i