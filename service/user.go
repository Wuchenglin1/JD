package service

import (
	"JD/dao"
	"JD/model"
	"JD/tool"
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

func SearchUserByPhone(phone string) (model.User, error) {
	return dao.SearchUserByPhone(phone)
}

// RegisterSendSMS 发送短信验证码
func RegisterSendSMS(phone string) error {
	return dao.SavePhoneVerifyCode(phone, "2022119")
}

// VerifyCodeByPhone 核对验证码是否正确
func VerifyCodeByPhone(u model.RegisterUser) (bool, error) {
	return dao.CheckVerifyCodeByPhone(u)
}

// SearchUserByEmail 通过email查找User
func SearchUserByEmail(email string) (model.RegisterUser, error) {
	iu, err := dao.SearchUserByEmail(email)
	u := model.RegisterUser{
		Id:       iu.Id,
		UserName: iu.UserName,
		Password: iu.Password,
		Phone:    iu.Phone,
		Email:    iu.Email,
	}
	return u, err
}

// RegisterSendEmail 发送注册时的邮箱验证码
func RegisterSendEmail(email string) error {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	config := tool.GetConfig()

	user := config.Email.User         //控制台创建的发信地址
	password := config.Email.Password //控制台设置的SMTP密码
	host := config.Email.Host         //smtpdm.aliyun.com:25
	to := []string{email}             //收件人地址（可以用,隔开添加多个账号群发消息）
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
	fmt.Println("正在给" + email + "发送短信，请稍等...")

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	contentType := "Content-Type: text/" + mailType + "; charset=UTF-8"
	msg := []byte("To: " + toAddress + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nReply-To: " + "\r\nCc: " + "\r\nBcc: " + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(host, auth, user, to, msg)

	err = dao.SaveEmailVerifyCode(email, code)

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

func SaveUser(u model.RegisterUser) error {
	return dao.SaveUser(u)
}

func SearchUserByUserName(userName string) (model.User, error) {
	return dao.SearchUserByUserName(userName)
}

func CheckUserByAccount(u model.User) (bool, error) {
	iu, err := dao.SearchUserByUserName(u.UserName)
	if err == nil {
		//密码错误
		if iu.Password != u.Password {
			return false, nil
		}
		return true, nil
	}
	iu, err = dao.SearchUserByEmail(u.UserName)
	if err == nil {
		//密码错误
		if iu.Password != u.Password {
			return false, nil
		}
		return true, nil
	}
	iu, err = dao.SearchUserByPhone(u.UserName)
	if err == nil {
		//密码错误
		if iu.Password != u.Password {
			return false, nil
		}
		return true, nil
	}
	return false, err
}
