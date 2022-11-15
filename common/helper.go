package common

import (
	"fmt"
	"gintest/config"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// GetConfig
// 获取配置文件
func GetConfig() *config.Config {
	yamlFile, err := ioutil.ReadFile("../config/application.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	var _config *config.Config

	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return _config
}

// SendEmail
// 发送邮件
func SendEmail(ToUserEmail string, code string) {

	// 获取配置文件信息
	_config := GetConfig()

	sender := _config.QQEmail.User                                                                    //发送者qq邮箱
	authCode := _config.QQEmail.AuthCode                                                              //qq邮箱授权码
	mailTitle := "FGC1024邮箱验证"                                                                        //邮件标题
	mailBody := "亲爱的" + ToUserEmail + "：\n\n欢迎注册FGC1024平台。您的验证码为 " + code + "\n\n如果您没有申请邮箱验证，请忽略此邮件。" //邮件内容,可以是html

	//接收者邮箱列表
	mailTo := []string{
		ToUserEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)       //发送者腾讯企业邮箱账号
	m.SetHeader("To", mailTo...)      //接收者邮箱列表
	m.SetHeader("Subject", mailTitle) //邮件标题
	m.SetBody("text/html", mailBody)  //邮件内容,可以是html

	//发送邮件服务器、端口、发送者qq邮箱、qq邮箱授权码
	//服务器地址和端口是腾讯的
	d := gomail.NewDialer(_config.QQEmail.Host, _config.QQEmail.Port, sender, authCode)
	if err := d.DialAndSend(m); err != nil {
		log.Println("send mail failed", err)
		return
	}

	log.Println("邮件发送成功……")
}

// GetUUID
// 生成一个唯一的UUID
func GetUUID() uuid.UUID {
	return uuid.Must(uuid.NewV4(), nil)
}
