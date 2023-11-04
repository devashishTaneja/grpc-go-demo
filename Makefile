gen-proto:
	@echo "==> Generating proto bindings..."
	@protoc proto/*.proto --go_out=pb --go-grpc_out=pb

run:
	go run main.go