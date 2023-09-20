gql:
	go run github.com/99designs/gqlgen generate

generate: generate-mocks gql

create-migration:
	go run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir db/migrations/scripts $(name)

migrate:
	go run cmd/cli/main.go db migrate

generate-mocks:
	go get go.uber.org/mock/mockgen/model
	go install go.uber.org/mock/mockgen@latest
	mockgen -destination=./mocks/mock_jwt_client.go -package=mocks github.com/bludot/tempmee/order/internal/services/jwt JWTServiceImpl
	mockgen -destination=./mocks/mock_order_repository.go -package=mocks github.com/bludot/tempmee/order/internal/db/repositories/order OrderRepositoryImpl
	mockgen -destination=./mocks/mock_order_service.go -package=mocks github.com/bludot/tempmee/order/internal/services/order OrderServiceImpl

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...