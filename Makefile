DB_URL=postgresql://root:secret@localhost:5432/pantopia?sslmode=disable

postgres:
	docker run --name postgres12 --network pantopia-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root pantopia
dropdb:
	docker exec -it postgres12 dropdb pantopia
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1
db_docs:
	dbdocs build doc/db.ml
db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml
sqlc:
	sqlc generate

test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/mbaxamb3/pantopia/db/sqlc Store
	
.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock db_docs db_schema