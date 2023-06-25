# ==============================================================================
# Go migrate postgresql

migrate_up:
	migrate -path migrations/ -database "postgresql://postgres:postgres@localhost:5432/audit?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations/ -database "postgresql://postgres:postgres@localhost:5432/audit?sslmode=disable" -verbose down
