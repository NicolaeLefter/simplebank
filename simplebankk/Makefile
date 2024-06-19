postgres:
	docker run --name postgres-one -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it postgres-one createdb --username=root --owner=root simple_bankk

dropdb:
	docker exec -it postgres-one dropdb simple_bankk

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bankk?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bankk?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY:postgres createdb dropdb migrateup migratedown