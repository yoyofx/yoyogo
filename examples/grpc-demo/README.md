## protoc 安装
```bash
wget https://github.com/google/protobuf/releases/download/v3.15.6/protobuf-all-3.15.6.zip
unzip protobuf-all-3.15.6.zip && cd protobuf-3.15.6/
./configure
make
make install

protoc --version
```

## go插件
```bash
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2
```

## 安装GRpc依赖
```bash
go get -u google.golang.org/grpc@v1.29.1
```

## .proto 代码生成
```bash
protoc --go_out=plugins=grpc:. ./proto/*.proto
```

## java 代码生成
```bash
protoc --java_out=./ XX.proto
```

## grpc调试
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

```bash
grpcurl -plaintext localhost:31127 list

# helloworld.Greeter
```
```bash
grpcurl -plaintext localhost:31127 list helloworld.Greeter

#helloworld.Greeter.SayHello
#helloworld.Greeter.SayList
#helloworld.Greeter.SayRecord
#helloworld.Greeter.SayRoute

```

### grpc调用
```bash
grpcurl -plaintext -d '{"name":"Go"}' localhost:31127 helloworld.Greeter.SayHello

#{
#  "message": "hello.world.at.server: Go"
#}

```

## grpc服务
```go
import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
	yrpc "github.com/yoyofx/yoyogo/grpc"
	"google.golang.org/grpc"
	pb "grpc-demo/proto/helloworld"
	"grpc-demo/services"
)

func main() {
    configuration := abstractions.NewConfigurationBuilder().
        AddEnvironment().
        AddYamlFile("config").Build()
    
    hosting := yrpc.NewHostBuilder().
        UseConfiguration(configuration).Configure(func(app *yrpc.ApplicationBuilder) {
            //app.AddUnaryServerInterceptor( logger.UnaryServerInterceptor() )
            //app.AddStreamServerInterceptor( logger.StreamServerInterceptor() )
            app.AddGrpcService(func( server *grpc.Server, ctx *yrpc.ServiceContext ) {
                pb.RegisterGreeterServer( server, services.NewGreeterServer() )
            })
        }).
        ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
    
        }).Build()
    
    hosting.Run()
}
```