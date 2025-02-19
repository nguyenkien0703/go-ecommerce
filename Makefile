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

up_by_one:
	@cmd /c "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one"

# create a new migration
#example: make create_migration name=00001_pre_go_acc_user_verify_9999
create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

# Lệnh migrate database sử dụng goose trên Windows
upse:
	@cmd /c "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) up"

downse:
	@cmd /c "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) down"

resetse:
	@cmd /c "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) reset"


sqlgen:
	sqlc generate

swag:
	if exist cmd\swag rmdir /s /q cmd\swag
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs
.PHONY: dev downse upse resetse docker_build docker_stop docker_up sqlgen create_migration up_by_one swag
.PHONY: air # Tương tự như nodemon bên Node.js, tự động reload server khi code thay đổi
