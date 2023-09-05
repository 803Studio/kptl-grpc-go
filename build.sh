protoc \
	--proto_path=../kptl-protos \
	--go_out=./pkg \
	--go-grpc_out=./pkg \
	../kptl-protos/*/*.proto
