GOOSE_DBSTRING=root:topsecret@tcp(127.0.0.1:3307)/shopdevgo
GOOSE_MIGRATION_DIR=sql/schema
GOOSE_DRIVER=mysql

##name app
APP_NAME = server

# Chạy ứng dụng
dev:
	go run ./cmd/$(APP_NAME)

run:
	docker compose up -d && go run ./cmd/$(APP_NAME)

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

# Lệnh migrate database sử dụng goose trên Windows
upse:
	@cmd /c "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) up"

downse:
	@cmd /c "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) down"

resetse:
	@cmd /c "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) reset"

.PHONY: run downse upse resetse
.PHONY: air # Tương tự như nodemon bên Node.js, tự động reload server khi code thay đổi
