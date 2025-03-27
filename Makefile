generate-auth:
	@echo "Cleaning old generated files..."
	rm -rf auth-microservice/proto/*
	@echo "Generating new code..."
		protoc --proto_path=proto \
    --go_out=gen-proto \
    --go_opt=module=github.com/IKHINtech/erp_microservice_go_grpc \
    --go-grpc_out=gen-proto \
    --go-grpc_opt=module=github.com/IKHINtech/erp_microservice_go_grpc \
    proto/auth/v1/auth.proto
	# protoc --proto_path=proto \
	# 	--go_out=auth-microservice/proto \
	# 	--go_opt=module=github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/proto \
	# 	--go-grpc_out=auth-microservice/proto \
	# 	--go-grpc_opt=module=github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/proto \
	# 	proto/auth/v1/auth.proto
	@echo "Generation completed!"

run-auth:
	cd auth-microservice && go run cmd/main.go
