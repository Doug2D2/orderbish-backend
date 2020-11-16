all: build run

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o orderbish main.go
	docker build --tag=orderbish .

run:
	docker-compose up orderbish db

clean:
	docker container stop orderbish-backend_orderbish_1 || docker container rm orderbish-backend_orderbish_1 || true
	docker container stop orderbish-backend_db_1 || docker container rm orderbish-backend_db_1 || true

.PHONY: build run clean
