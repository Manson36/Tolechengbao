package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lechengbao/datamodels"
	"github.com/lechengbao/datasource"
	"github.com/lechengbao/gen_id"
	"github.com/lechengbao/models"
	"github.com/lechengbao/repository"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//TODO：中间件保持用户登录状态,密码加盐，Ret

func Register(c *gin.Context) {
	//GET请求返回用户名，并将用户名写入注册用户名管理中
	if c.Request.Method == "GET" {
		username, _ := gen_id.GetInt64ID()
		nameStr := strconv.FormatInt(username, 10)
		//NameMsg = append(NameMsg, username)
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
			c.JSON(http.StatusOK, gin.H{
				"err":"参数解析错误",
			})
			return
		}

		nickName := strings.TrimSpace(body.Nickname)
		if nickName== "" {
			c.JSON(http.StatusOK, gin.H{
				"err":"请输入你的昵称",
			})
			return
		}

		password := strings.TrimSpace(body.Password)
		if password == "" {
			c.JSON(http.StatusOK, gin.H{
				"err":"请输入你的密码",
			})
			return
		}
		//redis库中查找注册用户名
		username := strings.TrimSpace(body.Username)
		_, err := datasource.Rds.Get(context.Background(), username).Result()
		if err == redis.Nil {
			c.JSON(http.StatusOK, gin.H{
				"err":"请返回正确的用户名",
			})
			return
		}else if err != nil {
			panic(err)
		}

		//这里省去了验证用户名在数据库中是否存在，因为用户名是自己生成的，不会重复
		id, _ := strconv.ParseInt(username, 10, 64)
		var user = datamodels.User{
			ID: id,
			Username:username,
			Nickname:nickName,
			Password:password,
		}
		if err := repository.CreateUser(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err":"数据库写入错误",
			})
			return
		}
		//删除redis中保存的用户名
		datasource.Rds.Del(context.Background(), username)

		c.JSON(http.StatusOK, gin.H{
			"msg":"注册成功",
			"user":user,
		})
	}
}

func Login(c *gin.Context) {
	var body models.LoginUserReqBody
	//请求参数绑定
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err":"用户登录参数解析失败",
		})
		fmt.Println(err)
		return
	}

	//验证用户登录请求信息
	username := strings.TrimSpace(body.Username)
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"err":"用户名为空",
		})
		return
	}
	password := strings.TrimSpace(body.Password)
	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"err":"密码为空",
		})
		return
	}

	user, err := repository.GetUser("username = ? ", username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err":"数据库用户信息获取失败",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"err":"用户名不存在",
		})
		return
	}
	if user.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"err":"密码错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Success":"登陆成功",
		"user":user,
	})
}
