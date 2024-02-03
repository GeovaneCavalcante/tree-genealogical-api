run:
	go run cmd/server.go
dev:
	~/go/bin/reflex -r "\.go" -s -- sh -c "go run cmd/server.go"