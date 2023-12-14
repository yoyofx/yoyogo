module websockethub

go 1.18

require (
	github.com/fasthttp/websocket v1.5.7
	github.com/yoyofx/yoyogo v0.0.0
	github.com/yoyofxteam/dependencyinjection v1.0.1
)

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace github.com/yoyofx/yoyogo => ../../
