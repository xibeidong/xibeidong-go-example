

 
### **注解：**

* @Summary 摘要

* @Produce API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等

* @Param 参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释

* @Success 响应成功，从左到右分别为：状态码、参数类型、数据类型、注释

* @Failure 响应失败，从左到右分别为：状态码、参数类型、数据类型、注释

* @Router 路由，从左到右分别为：路由地址，HTTP方法

### **安装**
```
go get -u github.com/swaggo/swag/cmd/swag
```
安装gin-swagger扩展
```
go get -u -v github.com/swaggo/gin-swagger
go get -u -v github.com/swaggo/files
go get -u -v github.com/alecthomas/template
```
### **运行**

进入main.go目录，执行
`$ swag init`
或者
`$ swag init main.go`
会自动生成docs

启动项目
访问 `localhost:8080/swagger/index.html`