package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lechengbao/datasource"
	"github.com/lechengbao/gen_id"
	"github.com/lechengbao/models"
	"github.com/lechengbao/services"
	"net/http"
	"strconv"
	"time"
)

//TODO：中间件保持用户登录状态,密码加盐，Ret

func Register(c *gin.Context) {
	//GET请求返回用户名，并将用户名写入注册用户名管理中
	if c.Request.Method == "GET" {
		username, _ := gen_id.GetInt64ID()
		nameStr := strconv.FormatInt(username, 10)
		//将用户名存入缓存中，注册完成删除
		datasource.Rds.Set(context.Background(), nameStr, nameStr, time.Second*3600)

		c.JSON(http.StatusOK, gin.H{
			"username":nameStr,
		})
	}

	//POST请求对用户进行注册
	if c.Request.Method == "POST" {
		var body models.AddUserReqBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusOK, &models.Ret{Code:400, Msg:"用户注册参数解析错误"})
			return
		}

		data := services.RegisterService(&body)

		c.JSON(http.StatusOK, data)
	}
}

func Login(c *gin.Context) {
	var body models.LoginUserReqBody
	//请求参数绑定
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusOK, &models.Ret{Code:400, Msg:"用户登录参数解析错误"})
		return
	}

	fmt.Println(body.Username)
	data := services.LoginService(&body)

	c.JSON(http.StatusOK, data)
}
