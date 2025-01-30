include ./.env
export

migrate:
	atlas migrate apply --url $(POSTGRES_DSN) --exec-order non-linear

server:
	go run cmd/server/main.go
