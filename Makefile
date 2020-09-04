

proto:
	protoc -I./demo --eva_out=plugins=grpc+eva:./demo demo.proto
	protoc -I./demo --grpc-gateway_out=logtostderr=true,paths=source_relative:./demo demo.proto

grpc:
	protoc -I./hello --go_out=plugins=grpc:./hello/demo demo.proto