SERVICE=link
PROTO_FILE=pb/$(SERVICE)/pb.proto

gen-proto:
	protoc -I/usr/local/include -I. \
        -I$(GOPATH)/bin \
        --proto_path=${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
        --validate_out="lang=go:." \
        --micro_out=. \
        --go_out=. \
        $(PROTO_FILE)

.PHONY: gen-proto