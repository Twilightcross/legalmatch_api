include .env

DB_URL := $(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)
GOOSE := goose -dir ./migrations mysql "$(DB_URL)"

migrate-up:
	$(GOOSE) up

migrate-down:
	$(GOOSE) down

migrate-status:
	$(GOOSE) status

migrate-reset:
	@echo "⚠️  모든 테이블이 삭제됩니다. 계속할까요? (y/n)"
	@read ans && [ "$$ans" = "y" ] && $(GOOSE) reset && $(GOOSE) up
