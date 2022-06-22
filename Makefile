SRC_DIR=.
DST_DIR=.

test:
	go test -v ./...

proto:
	protoc -I=${SRC_DIR} \
	--go_opt=paths=source_relative \
	--go_out=${DST_DIR} \
    --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	api/api.proto
