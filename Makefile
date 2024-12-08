generate_sso:
	@protoc -I ./proto sso.proto --go_out=./gen --go-grpc_out=./gen --go_opt=Muser.proto=github.com/DanKo-code/FitnessCenter-Protobuf/gen/FitnessCenter.protobuf.user

generate_user:
	@protoc -I ./proto user.proto --go_out=./gen --go-grpc_out=./gen

generate_coach:
	@protoc -I ./proto coach.proto --go_out=./gen --go-grpc_out=./gen

generate_service:
	@protoc -I ./proto service.proto --go_out=./gen --go-grpc_out=./gen

generate_abonement:
	@protoc -I ./proto abonement.proto --go_out=./gen --go-grpc_out=./gen

generate_review:
	@protoc -I ./proto review.proto --go_out=./gen --go-grpc_out=./gen

generate_all: generate_sso generate_user generate_coach generate_service generate_abonement generate_review
	@echo "All proto file have been generated"
