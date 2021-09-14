postgres:
	docker run --rm -d \
	 --name postgres12 \
	 -e POSTGRES_USER=root \
	 -e POSTGRES_PASSWORD=secret \
	 -v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
	 -p 5432:5432 postgres:12-alpine

postgres_cleanup:
	docker rm postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root nala

dropdb:
	docker exec -it postgres12 dropdb nala

migrateup:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/nala?sslmode=disable" -verbose up

migrateup_last:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/nala?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/nala?sslmode=disable" -verbose down

migratedown_last:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/nala?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./db/sqlc/

server:
	go run main.go

.PHONY: postgres postgres_cleanup createdb dropdb migrateup migrateup_last migratedown migratedown_last sqlc test server