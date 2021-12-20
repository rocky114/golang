pb:
	protoc --proto_path=./pb --go_out . --go-grpc_out . --grpc-gateway_out . ./pb/*.proto

image-server:
	docker build -t rocky114/server:0.1.2 -f build/server/Dockerfile .

image-client:
		docker build -t rocky114/client:0.1.1 -f build/client/Dockerfile .

