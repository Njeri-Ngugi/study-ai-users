include ./.env
export

hash:
	atlas migrate hash

diff:
	atlas migrate diff --env gorm $(migration_name)

migrate:
	atlas migrate apply --url $(POSTGRES_DSN) --exec-order non-linear

server:
	go run cmd/server/main.go

toolbox:
	go get github.com/Njeri-Ngugi/toolbox@latest