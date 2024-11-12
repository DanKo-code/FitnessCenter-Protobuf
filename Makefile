generate:
		@protoc -I ./proto sso.proto --go_out=./gen --go-grpc_out=./gen