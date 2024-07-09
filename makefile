# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

# MySQL command variables
MYSQL_CMD = mysql -u $(MYSQL_USER) -p$(MYSQL_PASSWORD) -h $(MYSQL_HOST) -P $(MYSQL_PORT) $(MYSQL_DATABASE)
MIGRATE_CMD = migrate -path=./work/migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)"

.PHONY: create-migrate
create-migrate:
	migrate create -ext sql -dir ./work/migrations -seq $(file-name)

.PHONY: migrate-up
migrate-up:
	$(MIGRATE_CMD) up

.PHONY: migrate-down
migrate-down:
	$(MIGRATE_CMD) down



