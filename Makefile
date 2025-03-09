dsn=postgres://coop_gardens:password@172.17.0.1:5432/coop_gardens?sslmode=disable

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