db2sql:
	dbml2sql --postgres -o db/doc/schema.sql db/doc/sql.dbml
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todoapp
dropdb:
	docker exec -it postgres12 dropdb todoapp
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateupALL:	
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todoapp?sslmode=disable" -verbose up
migratedownALL:	
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todoapp?sslmode=disable" -verbose down
sqlc:
	sqlc generate