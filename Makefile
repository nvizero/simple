postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

mysql:
	docker run --name mysql6 -p 3319:3306 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d mysql:5.7
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple

dropdb:
	docker exec -it postgres dropdb  --owner=root simple

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
