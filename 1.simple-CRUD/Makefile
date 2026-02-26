include .env
export

DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)&TimeZone=$(DB_TIMEZONE)

migrate-up:
	migrate.exe -path ./internal/db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate.exe -path ./internal/db/migrations -database "$(DB_URL)" down 1