project_name = myapp

run-dev:
	gow run app.go

requirements:
	go mod tidy
	go install github.com/mitranim/gow@latest

clean-packages:
	go clean -modcache


build:
	go build -a -o app .

run:
	./app -prod


auto-migrate:
	go run cmd/auto-migrate.go
