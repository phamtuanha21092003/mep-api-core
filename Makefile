.PHONY: proto, tusd-build

proto:
	@echo "Cleaning old generated files..."
	@rm -rf grpc/types
	@mkdir -p grpc/types
	@echo "Generating Go code from protobuf..."
	@protoc --proto_path=grpc \
		--go_out=grpc/types --go_opt=paths=source_relative \
		--go-grpc_out=grpc/types --go-grpc_opt=paths=source_relative \
		grpc/proto/*.proto
	@echo "Protobuf generation complete."


tusd-build:
	docker compose build tusd