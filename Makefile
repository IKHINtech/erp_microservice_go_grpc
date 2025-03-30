generate-auth:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative auth-microservice/pb/v1/auth.proto
	# @echo "Cleaning old generated files..."
	# rm -rf auth-microservice/gen-proto/proto/auth*
	# @echo "Generating new code..."
	# 	protoc --proto_path=proto \
	#    --go_out=gen-proto \
	#    --go_opt=module=github.com/IKHINtech/erp_microservice_go_grpc \
	#    --go-grpc_out=gen-proto \
	#    --go-grpc_opt=module=github.com/IKHINtech/erp_microservice_go_grpc \
	#    proto/auth/v1/auth.proto
	# protoc --proto_path=proto \
	# 	--go_out=auth-microservice/proto \
	# 	--go_opt=module=github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/proto \
	# 	--go-grpc_out=auth-microservice/proto \
	# 	--go-grpc_opt=module=github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/proto \
	# 	proto/auth/v1/auth.proto
	@echo "Generation completed!"

generate-organization:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative organization-microservice/pb/v1/organization.proto

run-auth:
	cd auth-microservice && go run cmd/main.go

run-api-gateway:
	cd api-gateway && go run cmd/main.go

run-organiation:
	cd organization-microservice && go run cmd/main.go

run-all:
	run-auth
	run-api-gateway

generate-docs:
	cd api-gateway && swag init -g cmd/main.go -o docs



