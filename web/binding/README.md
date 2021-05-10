## 参考Gin, 并与Gin Binding & Valvalidate API 保持一致.
```go
// 绑定为json
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// Example for binding JSON ({"user": "manu", "password": "123"})
err := ctx.Bind(Login)
```
```bash
curl -v -X POST \
  http://localhost:8080/loginJSON \
  -H 'content-type: application/json' \
  -d '{ "user": "manu" }'
```

##跳过验证：
当使用上面的curl命令运行上面的示例时，返回错误，因为示例中Password字段使用了binding:"required"，如果我们使用binding:"-"，那么它就不会报错。