protoc -I . --go_out=plugins=grpc:../proto  product.proto
https://github.com/grpc/grpc-go/issues/3347
如果还不行 + go mod vendor