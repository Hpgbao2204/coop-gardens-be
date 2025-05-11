include .env

dsn=${GOOSE_DBSTRING}
migrationDir=internal/db/migrations

up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} up
down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} down
status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} status
up-by-one:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} up-by-one
reset:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} reset
run:
	go run cmd/main.go