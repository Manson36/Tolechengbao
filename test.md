##用户注册：
http://localhost:8080/user/register

|字段|说明|类型|
|---|---|---|
|username|用户名|string|
|nickname|昵称|string|
|password|密码|string|

http请求方式：GET
返回值
```
{
  "username": "176565518336"
}
```
http请求方式：POST
请求内容
```
{
   "username":"176565518336",
   "nickname":"吉吉蜥蜴",
   "password":"12345"
   }
```
注册成功返回json数据包
```$xslt
{
   "Code": 200,
    "Msg": "Success",
    "Data": {
      "id": 176565518336,
      "username": "176565518336",
      "nickname": "吉吉蜥蜴",
      "password": "12345",
      "createAt": "2020-09-25T22:47:34.952436+08:00",
      "updateAt": "2020-09-25T22:47:34.952436+08:00",
      "removedAt": null,
      "removed": false,
      "disabled": false,
      "disabledAt": null
    }
}
```
密码为空返回json数据包
```$xslt
{
  "Code": 400,
    "Msg": "请输入密码",
    "Data": null
}
```

##用户登录
请求方法：POST
http://localhost:8080/user/login

|字段|说明|类型|
|---|---|---|
|username|用户名|string|
|password|密码|string|

请求内容
```
{
   {
   "username":"176565518336",
   "password":"12345"
   }
```

登录成功返回json数据包
```$xslt
{
   "Code": 200,
     "Msg": "登录成功",
     "Data": {
       "id": 176565518336,
       "username": "176565518336",
       "nickname": "吉吉蜥蜴",
       "password": "12345",
       "createAt": "2020-09-25T22:47:34.952436+08:00",
       "updateAt": "2020-09-25T22:47:34.952436+08:00",
       "removedAt": null,
       "removed": false,
       "disabled": false,
       "disabledAt": null
     }
}
```

密码错误返回json数据包
```$xslt
{
     "Code": 400,
     "Msg": "密码错误，请重新输入",
     "Data": null
}
```

用户名错误返回json数据包
```$xslt
{
     "Code": 400,
     "Msg": "用户名不存在",
     "Data": null
}
```