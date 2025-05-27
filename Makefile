include .env
export $(shell sed 's/=.*//' .env)

MIGRATE_DIR=cmd/migrate
CONTAINER_NAME=immin_infra_mariadb.1.2eo50qxw0r3hg1v3qcu8n0le8

migrate:
	@echo "Running all *_up.sql files in order inside container..."
	@for f in $(shell find $(MIGRATE_DIR) -name '*_up.sql' | sort); do \
		echo "Applying $$f..."; \
		docker exec -i $(CONTAINER_NAME) mariadb -u$(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) < $$f || exit 1; \
	done
	@echo "Migration complete."


rollback:
	@echo "Rolling back all *_down.sql files in reverse order inside container..."
	@for f in $(shell find $(MIGRATE_DIR) -name '*_down.sql' | sort); do \
		echo "Applying $$f..."; \
		docker exec -i $(CONTAINER_NAME) mariadb -u$(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) < $$f || exit 1; \
	done
	@echo "Rollback complete."
