export MYSQL_URL='mysql://root:2secret@tcp(localhost:3307)/db-simple-forum'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down

gorun:
	@ go run cmd/main.go