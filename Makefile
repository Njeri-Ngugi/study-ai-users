migrate:
	atlas migrate apply --url "postgres://$(user):$(pwd)@localhost:5432/study_assistant?sslmode=disable&search_path=public" --exec-order non-linear

server:
	go run cmd/server/main.go
