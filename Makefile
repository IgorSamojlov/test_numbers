.PHONY: run vet
dev_run:
	air --build.cmd "go build -C cmd/http -o ../../build/http" --build.entrypoint "./build/http"
vet:
	go vet ./...
db_prepare:
	createdb -U igor dev_numbers
	psql -U igor -d dev_numbers -c "CREATE TABLE IF NOT EXISTS numbers (id SERIAL PRIMARY KEY, value INTEGER NOT NULL, created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW());"
db_drop:
	psql -U igor -c "DROP DATABASE dev_numbers;"
run:
	docker compose build
	docker compose up
