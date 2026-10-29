include ./.env.migration

DB_URL=postgres://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable

MIGRATION_PATH=db/migrations
SEEDER_PATH=db/seeds


migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq create_$(NAME)_table


migrate-up:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) up


migrate-down:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) down

migrate-version:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) version

print-db-url:
	@echo $(DB_URL)

seed:
	@psql $(DB_URL) < $(SEEDER_PATH)/001_roles.sql
	@psql $(DB_URL) < $(SEEDER_PATH)/002_event_categories.sql
	@psql $(DB_URL) < $(SEEDER_PATH)/003_community_categories.sql
	@psql $(DB_URL) < $(SEEDER_PATH)/004_communities_events.sql

.PHONY: migrate-docker-up migrate-docker-version

migrate-docker-up:
	@docker compose run --rm migrate up

migrate-docker-version:
	@docker compose run --rm migrate version