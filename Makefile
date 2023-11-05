server-gen-proto:
	@echo "==> Generating proto bindings..."
	@protoc server/proto/*.proto --go_out=server/pb --go-grpc_out=server/pb

client-gen-proto:
	@echo "==> Generating proto bindings..."
	@protoc client/proto/*.proto --go_out=client/pb --go-grpc_out=client/pb

grpc-server:
	@echo "==> Starting server..."
	@cd server && go run main.go && cd ..

grpc-client:
	@echo "==> Starting client..."
	@cd client && go run main.go && cd ..