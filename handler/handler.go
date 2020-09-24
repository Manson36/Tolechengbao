package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lechengbao/gen_id"
	"net/http"
)

//这个放弃了，新写了在 user.go 中

type UserInfo struct {
	Username int64 `form:"username"`
	Password string `form:"password"`
}

var UserMessage = make([]UserInfo, 0)

func RegisterHandler(c *gin.Context) {
	var u UserInfo
	u.Username, _ = gen_id.GetInt64ID()

	if c.Request.Method == "POST" {
		if err := c.ShouldBind(&u); err != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"err":"用户名密码不能为空",
			})
			return
		}

		UserMessage = append(UserMessage, u)
		c.HTML(http.StatusOK, "home.html", nil)
	} else {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"username":u.Username,
		})
	}
}

func LoginHandler(c *gin.Context) {
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