.PHONY: protoItem

protoItem: protoc --go_out=./_proto --go_opt=paths=source_relative --go-grpc_out=./_proto --go-grpc_opt=paths=source_relative ./api/item.proto