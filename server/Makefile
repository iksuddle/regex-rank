build:
	go build -o ./bin/rgx ./cmd/regex-rank/

run: build
	./bin/rgx

migration-up:
	go run ./cmd/migrate/ up

migration-down:
	go run ./cmd/migrate/ down
