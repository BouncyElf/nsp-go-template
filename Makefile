gen:
	protoc --go_out=./pb/${pb_path} --go_opt=paths=source_relative --go-grpc_out=./pb/${pb_path} --go-grpc_opt=paths=source_relative ./pb/idl/*.proto -I ./pb/idl

test:
	go test ./...

run:
	go run . -c ./config/config.toml
