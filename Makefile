DB_URL=postgresql://root:root@localhost:5432/simple?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14-alpine

mysql:
	docker run --name mysql6 -p 3319:3306 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d mysql:5.7
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple

dropdb:
	docker exec -it postgres dropdb  --owner=root simple

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination mockdb/store.go simple/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
