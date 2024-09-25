build:
	go build -o ./bin/rgx ./cmd/api/

run: build
	./bin/rgx

# here filter-out removes `migration` from the arguments following the make command
migration:
	migrate create -ext sql -dir ./cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run ./cmd/migrate/main.go up

migrate-down:
	go run ./cmd/migrate/main.go down

create: problems.toml
	@go run ./cmd/create/ problems.toml
