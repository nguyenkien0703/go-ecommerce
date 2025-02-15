##name app
#APP_NAME = server
#run:
#	go run ./cmd/${APP_NAME}/





##name app
APP_NAME = server

# chay ung dung
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

.PHONY: run
.PHONY: air # tuong tu nhu nodemon ben nodejs ay, tu dong reload server khi code change