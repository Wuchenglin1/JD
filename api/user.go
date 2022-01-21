package api

import (
	"JD/model"
	"JD/service"
	"JD/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// RegisterSendSMS 注册检验手机号和发送验证码
func RegisterSendSMS(c *gin.Context) {
	u := model.User{}
	u.Phone = c.PostForm("phone")
	if u.Phone == "" {
		tool.RespErrWithData(c, false, "手机号不可为空")
		return
	}
	if len(u.Phone) != 11 {
		tool.RespErrWithData(c, false, "手机号不合法")
		return
	}
	fmt.Println(u.Phone)
	_, err := service.SearchUserByPhone(u.Phone)
	if err == nil {
		tool.RespErrWithData(c, false, "手机号已被注册")
		return
	}
	err = service.RegisterSendSMS(u.Phone)
	if err != nil {
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccess(c)
}

// CheckRegisterSMS 检查短信验证码是否正确
func CheckRegisterSMS(c *gin.Context) {
	u := model.RegisterUser{}
	u.Phone = c.PostForm("phone")
	u.VerifyCode = c.PostForm("verifyCode")
	if u.VerifyCode == "" {
		tool.RespErrWithData(c, false, "验证码不能为空")
		return
	}
	IsCorrect, err := service.VerifyCodeByPhone(u)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "未发送验证码")
		return
	}
	if !IsCorrect {
		tool.RespErrWithData(c, false, "验证码错误")
		return
	}
	tool.RespSuccess(c)
}

// RegisterSendEmail 发送邮箱验证码
func RegisterSendEmail(c *gin.Context) {
	u := model.RegisterUser{}
	u.Email = c.PostForm("email")
	if u.Email == "" {
		tool.RespErrWithData(c, false, "邮箱不能为空")
		return
	}
	n := strings.Index(u.Email, "@")
	if n == -1 {
		tool.RespErrWithData(c, false, "邮箱填写错误")
		return
	}
	_, err := service.SearchUserByEmail(u.Email)
	if err == nil {
		tool.RespErrWithData(c, false, "邮箱已被注册")
		return
	}
	err = service.RegisterSendEmail(u.Email)
	if err != nil {
		tool.RespErrWithData(c, false, "服务器错误")
	}
}

// Register 检验所有信息的正误
func Register(c *gin.Context) {
	u := model.RegisterUser{}
	u.UserName = c.PostForm("userName")
	u.Password = c.PostForm("password")
	u.Email = c.PostForm("email")
	u.Phone = c.PostForm("phone")
	if u.UserName == "" {
		tool.RespErrWithData(c, false, "用户名不能为空")
		return
	}
	if len(u.UserName) >= 20 {
		tool.RespErrWithData(c, false, "用户名太长了")
		return
	}
	_, err := service.SearchUserByUserName(u.UserName)
	if err == nil {
		tool.RespErrWithData(c, false, "该用户名已被使用，请更换名称")
	}
	if len(u.Password) <= 6 {
		tool.RespErrWithData(c, false, "密码不能小于6个字符")
		return
	}
	if len(u.Password) >= 16 {
		tool.RespErrWithData(c, false, "密码不能大于16个字符")
		return
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashPassword)
	if u.Email == "" {
		tool.RespErrWithData(c, false, "邮箱不能为空")
		return
	}
	n := strings.Index(u.Email, "@")
	if n == -1 {
		tool.RespErrWithData(c, false, "邮箱填写错误")
		return
	}
	_, err = service.SearchUserByEmail(u.Email)
	if err == nil {
		tool.RespErrWithData(c, false, "邮箱已被注册")
		return
	}
	u.VerifyCode = c.PostForm("verifyCode")

	if u.VerifyCode == "" {
		tool.RespErrWithData(c, false, "请输入验证码")
		return
	}
	IsCorrect, err2 := service.CheckVerifyCodeByEmail(u)
	if err2 != nil {
		fmt.Println("err2", err2)
		tool.RespErrWithData(c, false, "未发送验证码")
		return
	}
	if !IsCorrect {
		tool.RespErrWithData(c, false, "验证码错误")
		return
	}
	err = service.SaveUser(u)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccess(c)
}

func Login(c *gin.Context) {
	u := model.User{}
	u.UserName = c.PostForm("account")
	u.Password = c.PostForm("password")
	if u.UserName == "" && u.Password == "" {
		tool.RespErrWithData(c, false, "请输入账户名和密码")
		return
	}
	if u.UserName == "" {
		tool.RespErrWithData(c, false, "请输入账户名")
		return
	}
	if u.Password == "" {
		tool.RespErrWithData(c, false, "请输入密码")
		return
	}
	is, err := service.CheckUserByAccount(u)
	if err != nil {
		tool.RespErrWithData(c, false, "账号不存在")
		return
	}
	if !is {
		tool.RespErrWithData(c, false, "账户名与密码不匹配，请重新输入")
		return
	}
	//创建一个存在一天的refreshToken
	rt, err := service.CreateToken(u, 86400, "refreshToken")
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	//创建一个存在5min的token
	t, err := service.CreateToken(u, 300, "token")
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	c.JSON(200, gin.H{
		"status":       true,
		"data":         "登录成功",
		"token":        t,
		"refreshToken": rt,
	})
}
