postgres:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=ROOT  -e POSTGRES_PASSWORD=secret -d postgres:17-alpine3.20

createdb:
	sudo docker exec -it postgres17 createdb --username=ROOT --owner=ROOT mysimple-bank

dropdb:
	sudo docker exec -it postgres17 dropdb --username=ROOT mysimple-bank

migrateup:
	migrate -path db/migration -database "postgresql://ROOT:secret@localhost:5432/mysimple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://ROOT:secret@localhost:5432/mysimple-bank?sslmode=disable" -verbose down

start:
	sudo docker start postgres17

stop:
	sudo docker stop postgres17

list:
	docker ps

rmcontainer:
	sudo docker rm postgres17

sqlc:
	sqlc generate

test:
	go test -v -cover -count=1 ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown start stop list rmcontainer sqlc test server

## https://sqlc.dev/