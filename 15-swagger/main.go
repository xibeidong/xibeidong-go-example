package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "xibeidong-go-example/15-swagger/docs" // 记得docs导入进来
)

type UserInfo struct {
	Id   string
	Name string
	Age  int
}

// @title ToDoList API
// @version 0.0.1
// @description This is a sample Server pets
// @name FanOne
// @BasePath /api/v1
func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/v1/userinfo", GetUserInfo)
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

// GetUserInfo
// @Summary 获取user详细
// @Produce json
// @Param user_id body string false "标签ID"
// @Success 200 {object} UserInfo "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/v1/userinfo/{id} [get]
func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "2232",
		"name": "lucy",
		"age":  22,
	})
}
