protoc \
	--proto_path=/home/cdl/GolandProjects/kptl-protos \
	--go_out=./pkg \
	--go-grpc_out=./pkg \
	/home/cdl/GolandProjects/kptl-protos/*/*.proto
