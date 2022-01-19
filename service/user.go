package service

import (
	"JD/dao"
	"JD/model"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

func SearchUserByPhone(phone string) (model.User, error) {
	return dao.SearchUserByPhone(phone)
}

// RegisterSendSMS 发送验证码
func RegisterSendSMS(phone string) error {

}

// VerifyCodeByPhone 核对验证码是否正确
func VerifyCodeByPhone(u model.RegisterUser) (bool, error) {
	return dao.CheckVerifyCodeByEmail(u)
}

// SearchUserByEmail 通过email查找User
func SearchUserByEmail(email string) (model.RegisterUser, error) {
	return dao.SearchUserByEmail(email)
}

// RegisterSendEmail 发送注册时的邮箱验证码
func RegisterSendEmail(email string) error {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	user := "1336636739@wuchenglin.plus" //控制台创建的发信地址
	password := "WUchenglin2021214174"   //控制台设置的SMTP密码
	host := "smtpdm.aliyun.com:25"       //smtpdm.aliyun.com:25
	to := []string{email}                //收件人地址（可以用,隔开添加多个账号群发消息）
	toAddress := strings.Join(to, ";")
	subject := "破小站发送的验证码" //标题
	mailType := "html"     //发送类型
	body :=
		`
	<html>
	<body>
    <h4>验证码</h4>
    <p>您好，您的验证码为:` + code + `有效期为30min，请不要泄露给他人！</p>
	</body>
	</html>
	`
	fmt.Println("发送中 请稍等...")

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	contentType := "Content-Type: text/" + mailType + "; charset=UTF-8"
	msg := []byte("To: " + toAddress + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nReply-To: " + "\r\nCc: " + "\r\nBcc: " + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(host, auth, user, to, msg)

	err = dao.SaveEmailVerifyCode(email, code)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err != nil {
		fmt.Println("发送错误！")
		fmt.Println(err)
		return err
	} else {
		fmt.Println("发送成功！")
		return nil
	}
}

func CheckVerifyCodeByEmail(u model.RegisterUser) (bool, error) {
	return dao.CheckVerifyCodeByEmail(u)

}
