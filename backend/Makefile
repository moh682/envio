name?=missing_name

dev:
	air
 
migrate:
	go run cmd/migration/main.go

sqlc:
	sqlc generate

migration-file:
	migrate create -ext sql -dir internal/frameworks/sqlite/migrations -seq $(name)

tests:
	go test -v ./...

