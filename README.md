# GINTEST

## 第三方插件
- swagger `go get -u github.com/swaggo/swag`
  [主页链接](https://github.com/swaggo/gin-swagger)
```shell
swag init
```

```go 
import (
docs "github.com/go-project-name/docs"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example 接口名称
// @Schemes 表结构
// @Description do ping  描述
// @Tags example 标签
// @Accept json
// @Produce json
// @Param page query int false "page" 参数
// @Param size query int false "size" 参数
// @Param keyword query string false "trip_number"
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
```


- viper `go get github.com/spf13/viper`
- gomail `go get gopkg.in/gomail.v2`

```go
sender := _config.QQEmail.User       //发送者qq邮箱
authCode := _config.QQEmail.AuthCode //qq邮箱授权码
host := _config.QQEmail.Host
port := _config.QQEmail.Port

mailTitle := "邮件标题" //邮件标题
mailBody := "邮件内容"  //邮件内容,可以是html

//接收者邮箱列表
mailTo := []string{
    "464773078@qq.com",
}

m := gomail.NewMessage()
m.SetHeader("From", sender)       //发送者腾讯企业邮箱账号
m.SetHeader("To", mailTo...)      //接收者邮箱列表
m.SetHeader("Subject", mailTitle) //邮件标题
m.SetBody("text/html", mailBody)  //邮件内容,可以是html

//发送邮件服务器、端口、发送者qq邮箱、qq邮箱授权码
//服务器地址和端口是腾讯的
d := gomail.NewDialer(host, port, sender, authCode)
if err := d.DialAndSend(m); err != nil {
    log.Println("send mail failed", err)
    return
}

log.Println("success")
```