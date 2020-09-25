package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lechengbao/datamodels"
	"github.com/lechengbao/datasource"
	"github.com/lechengbao/models"
	"github.com/lechengbao/repository"
	"strconv"
	"strings"
)

func RegisterService(body *models.AddUserReqBody) *models.Ret {
	nickName := strings.TrimSpace(body.Nickname)
	if nickName== "" {
		return &models.Ret{Code:400, Msg:"请输入昵称"}
	}

	password := strings.TrimSpace(body.Password)
	if password == "" {
		return &models.Ret{Code:400, Msg:"请输入密码"}
	}

	//redis库中查找注册用户名
	username := strings.TrimSpace(body.Username)
	_, err := datasource.Rds.Get(context.Background(), username).Result()
	if err == redis.Nil {
		return &models.Ret{Code:400, Msg:"请不要修改用户名"}
	}else if err != nil {
		panic(err)
	}

	//这里省去了验证用户名在数据库中是否存在，逻辑与登录相似，user != nil, return
	id, _ := strconv.ParseInt(username, 10, 64)
	var user = datamodels.User{
		ID: id,
		Username:username,
		Nickname:nickName,
		Password:password,
	}
	if err := repository.CreateUser(&user); err != nil {
		return &models.Ret{Code:500, Msg:"用户信息入库失败"}
	}
	//删除redis中保存的用户名
	datasource.Rds.Del(context.Background(), username)

	return &models.Ret{Code:200, Msg:"Success", Data:user}
}

func LoginService(body *models.LoginUserReqBody) *models.Ret {
	//验证用户登录请求信息
	username := strings.TrimSpace(body.Username)
	if username == "" {
		return &models.Ret{Code:400, Msg:"请输入用户名"}
	}

	password := strings.TrimSpace(body.Password)
	if password == "" {
		return &models.Ret{Code:400, Msg:"请输入密码"}
	}

	user, err := repository.GetUser("username = ? ", username)
	if err != nil {
		return &models.Ret{Code:500, Msg:"数据库获取用户信息错误"}
	}

	if user == nil {
		return &models.Ret{Code:400, Msg:"用户名不存在"}
	}

	if user.Password != password {
		return &models.Ret{Code:400, Msg:"密码错误，请重新输入"}
	}

	return &models.Ret{Code:200, Msg:"登录成功", Data:user}
}