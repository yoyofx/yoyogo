module simpleweb

go 1.16

require (
	github.com/fasthttp/websocket v1.4.3
	github.com/go-sql-driver/mysql v1.6.0
	github.com/yoyofx/yoyogo v0.0.0
	gorm.io/gorm v1.21.11
)

replace github.com/yoyofx/yoyogo => ../../
