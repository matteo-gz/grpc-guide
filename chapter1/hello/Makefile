pb:  # generate pb.go
	protoc  \
	--proto_path=.  \
	--go_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative:. \
	proto/greeter.proto


srv:
	go run cmd/server/server.go

call:
	go run cmd/client/client.go