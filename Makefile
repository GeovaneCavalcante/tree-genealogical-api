run:
	go run cmd/server.go
	
dev:
	~/go/bin/reflex -r "\.go" -s -- sh -c "go run cmd/server.go"

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	~/go/bin/swag init -g ./cmd/server.go --output ./docs

build-mocks:
	go install go.uber.org/mock/mockgen@latest
	~/go/bin/mockgen -source=familytree/familytree.go -destination=familytree/mock/familytree.go
	~/go/bin/mockgen -source=person/person.go -destination=person/mock/person.go
	~/go/bin/mockgen -source=relationship/relationship.go -destination=relationship/mock/relationship.go
	
test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html


docker-build-image:
	docker build -t geovanedeveloper/tree-genealogical-api:latest -f Dockerfile.prod .

docker-run:
	docker run --rm -p 8080:8080 geovanedeveloper/tree-genealogical-api:latest
