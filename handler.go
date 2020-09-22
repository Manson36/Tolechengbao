package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id int
	Username string `form:"username"`
	Password string `form:"password"`
}

var UserMessage = make([]UserInfo, 0)

func registerHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		var u UserInfo
		if err := c.ShouldBind(&u); err != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"err":"用户名密码不能为空",
			})
			return
		}

		for _, user := range UserMessage {
			if u.Username == user.Username {
				c.HTML(http.StatusOK, "register.html", gin.H{
					"err":"用户名已经存在",
				})
				return
			}
		}

		u.Id = GenId()
		UserMessage = append(UserMessage, u)
		c.HTML(http.StatusOK, "home.html", nil)
	} else {
		c.HTML(http.StatusOK, "register.html", nil)
	}
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		var u UserInfo
		if err := c.ShouldBind(&u); err != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"err":"用户名密码不能为空",
			})
			return
		}

		for _, user := range UserMessage {
			if u.Username == user.Username {
				if u.Password == user.Password {
					c.HTML(http.StatusOK, "home.html", nil)
				} else {
					c.HTML(http.StatusOK, "login.html", gin.H{
						"err":"登录密码错误",
					})
				}
			} else {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"err":"用户不存在",
				})
			}
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}