module simpleweb

go 1.16

require (
	github.com/fasthttp/websocket v1.4.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/yoyofx/yoyogo v0.0.0
	gorm.io/gorm v1.20.12
)

replace github.com/yoyofx/yoyogo => ../../

//require github.com/yoyofx/yoyogo v1.5.5
