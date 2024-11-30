generate_sso:
	@protoc -I ./proto sso.proto --go_out=./gen --go-grpc_out=./gen

generate_user:
	@protoc -I ./proto user.proto --go_out=./gen --go-grpc_out=./gen

generate_coach:
	@protoc -I ./proto coach.proto --go_out=./gen --go-grpc_out=./gen

generate_service:
	@protoc -I ./proto service.proto --go_out=./gen --go-grpc_out=./gen