GOOSE_DBSTRING=root:topsecret@tcp(127.0.0.1:3307)/shopdevgo
GOOSE_MIGRATION_DIR=sql/schema
GOOSE_DRIVER=mysql

##name app
APP_NAME := server

# Chạy ứng dụng
docker_build:
	docker-compose up -d --build
	docker-compose ps

docker_stop:
	docker-compose stop
dev:
	go run ./cmd/$(APP_NAME)



docker_up:
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

.PHONY: dev downse upse resetse docker_build docker_stop docker_up
.PHONY: air # Tương tự như nodemon bên Node.js, tự động reload server khi code thay đổi
