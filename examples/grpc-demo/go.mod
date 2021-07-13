module grpc-demo

go 1.16

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/yoyofxteam/dependencyinjection v1.0.0
	github.com/yoyofx/yoyogo v0.0.0
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/yoyofx/yoyogo => ../../
