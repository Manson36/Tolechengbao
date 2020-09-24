##用户注册：
http://localhost:8080/user/register

|字段|说明|类型|
|---|---|---|
|username|用户名|int64|
|nickname|昵称|string|
|password|密码|string|

http请求方式：GET
返回值
```
{
  "username": 55569285120
}
```
http请求方式：POST
请求内容
```
{
   "username":55569285120,
   "nickname":"Kitty",
   "password":"123"
   }
```
注册成功返回json数据包
```$xslt
{
   "msg": "注册成功",
   "user": {
     "username": 55569285120,
     "nickname": "Kitty",
     "password": "123"
   }
}
```
密码为空返回json数据包
```$xslt
{
  "err": "请输入你的密码"
}
```

##用户登录
请求方法：POST
http://localhost:8080/user/login

|字段|说明|类型|
|---|---|---|
|username|用户名|int64|
|password|密码|string|

请求内容
```
{
   {
   "username":55569285120,
   "password":"123"
   }
```

登录成功返回json数据包
```$xslt
{
   "msg": "登录成功",
    "user": {
      "username": 55569285120,
      "nickname": "Kitty",
      "password": "123"
    }
}
```

密码错误返回json数据包
```$xslt
{
    "msg": "密码错误"
}
```