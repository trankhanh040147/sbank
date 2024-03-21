createdb:
	docker exec -it postgres12 createdb --username=root --owner=root sbank
dropdb:
	docker exec -it postgres12 dropdb sbank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/sbank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/sbank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc