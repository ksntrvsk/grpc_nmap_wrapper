build: 
	go build cmd/server/main.go
test:
	go test pkg/server/server_test.go
lint: 
	golangci-lint run
server:
	go run cmd/server/main.go
client:
	go run cmd/client/main.go