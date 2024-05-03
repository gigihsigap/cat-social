migrate-up:
	migrate -path db/migrations/ -database "postgresql://gap@localhost:5432/mydatabase?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations/ -database "postgresql://gap@localhost:5432/mydatabase?sslmode=disable" -verbose down