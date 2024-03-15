

- install golang
- install docker
- install protoc  // brew install protobuf
- install buf // brew install bufbuild/buf/buf

TODO: setup config for proto buf

Go follows a convention where source files are all lower case with underscore separating multiple words.



// Progo install 

Run go mod tidy to resolve the versions. Install by running

$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
This will place four binaries in your $GOBIN;

protoc-gen-grpc-gateway
protoc-gen-openapiv2
protoc-gen-go
protoc-gen-go-grpc