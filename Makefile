run:
	go run cmd/server.go
dev:
	~/go/bin/reflex -r "\.go" -s -- sh -c "go run cmd/server.go"

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	~/go/bin/swag init -g ./cmd/server.go --output ./docs