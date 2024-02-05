run:
	go run cmd/server.go
dev:
	~/go/bin/reflex -r "\.go" -s -- sh -c "go run cmd/server.go"

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	~/go/bin/swag init -g ./cmd/server.go --output ./docs

build-mocks:
	go install go.uber.org/mock/mockgen@latest
	
test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html