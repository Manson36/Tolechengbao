package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lechengbao/gen_id"
	"net/http"
	"sync"
)

type User struct {
	Username int64 `json:"username, string"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

var(
	mu = sync.Mutex{}
	UserMsg = make([]User, 0)
	NameMsg = make([]int64, 0) //解决GET请求注册页面时给的用户名和POST请求注册 用户名不一致问题
)

//TODO：UserMsg写入数据库中；NameMsg写入Redis中；中间件保持用户登录状态

func Register(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	//GET请求返回用户名，并将用户名写入注册用户名管理中
	if c.Request.Method == "GET" {
		username, _ := gen_id.GetInt64ID()
		NameMsg = append(NameMsg, username)

		c.JSON(http.StatusOK, gin.H{
			"username":username,
		})
	}

	//POST请求对用户进行注册
	if c.Request.Method == "POST" {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err":"参数解析错误",
			})
		}

		if u.Nickname == "" {
			c.JSON(http.StatusOK, gin.H{
				"err":"请输入你的昵称",
			})
			return
		}

		if u.Password == "" {
			c.JSON(http.StatusOK, gin.H{
				"err":"请输入你的密码",
			})
			return
		}

		//用户名管理中需要有申请的用户名
		for i, username := range NameMsg {
			if u.Username == username {
				UserMsg = append(UserMsg, u)	//将用户信息写入库中
				NameMsg = append(NameMsg[:i], NameMsg[i+1:]...)	//用户已经注册，删除用户名管理中该名

				c.JSON(http.StatusOK, gin.H{
					"msg":"注册成功",
					"user":u,
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"err":"请返回正确的用户名",
		})
	}
}

func Login(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var u User
	//请求参数绑定
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err":"用户登录参数解析失败",
		})
		fmt.Println(err)
	}

	//请求参数验证
	for _, user := range UserMsg {
		if user.Username == u.Username {
			if user.Password == u.Password {
				c.JSON(http.StatusOK, gin.H{
					"msg":"登录成功",
					"user":user,
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg":"密码错误",
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":"用户名不存在",
	})
}
