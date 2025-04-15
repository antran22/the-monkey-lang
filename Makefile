.PHONY: repl test
test:
	go test ./...

repl:
	go run ./main.go
