generate_sso:
	@protoc -I ./proto sso.proto --go_out=./gen --go-grpc_out=./gen

generate_user:
	@protoc -I ./proto user.proto --go_out=./gen --go-grpc_out=./gen