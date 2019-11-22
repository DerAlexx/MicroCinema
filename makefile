# This makefile will build all needed .proto, .. files.

proto:
	go install -v github.com/gogo/protobuf/protoc-gen-gogoslick
		protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
			--gogoslick_out=plugins=grpc:. users.proto