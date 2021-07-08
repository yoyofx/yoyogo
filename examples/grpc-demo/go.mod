module grpc-demo

go 1.16

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/yoyofx/yoyogo v0.0.0
	golang.org/x/exp v0.0.0-20200331195152-e8c3332aa8e5 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/yoyofx/yoyogo => ../../
