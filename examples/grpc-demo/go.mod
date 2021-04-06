module grpc-demo

go 1.16

require (
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/yoyofx/yoyogo v0.0.0
	golang.org/x/net v0.0.0-20210326220855-61e056675ecf // indirect
	golang.org/x/sys v0.0.0-20210326220804-49726bf1d181 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.26.0
)

replace github.com/yoyofx/yoyogo => ../../
