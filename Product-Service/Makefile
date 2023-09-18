build:
	@ cd bin; go build ../cmd/main.go; echo "Build Done😊";

run: build
	@ ./bin/main;

proto-gen:
	@  echo "[INFO]: generating protofiles. "; protoc --go_out=proto \
  		--go-grpc_out=proto proto/product_service.proto

grpc-gatewy-gen: proto-gen
	@ echo "[INFO]: generating grpc-gateway files. ";
	@ protoc --grpc-gateway_out . \
         --grpc-gateway_opt logtostderr=true \
         --grpc-gateway_opt paths=source_relative \
         ./proto/product_service.proto
