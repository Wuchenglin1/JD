package api

import (
	"JD/model"
	"JD/service"
	"JD/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// RegisterSendSMS 注册界面需要先检验手机号
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
	_, err := service.SearchUserByPhone(u.Phone)
	if err == nil {
		tool.RespErrWithData(c, false, "手机号已被注册")
		return
	}
	err = service.RegisterSendSMS(u.Phone) //这里的验证码要存储到RegisterUser表中
	if err != nil {
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespErrWithData(c, true, "")
}

// CheckRegisterSMS 检查验证码是否正确
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
	tool.RespErrWithData(c, true, "")
}

// Register 注册第二步，账号密码，以及邮箱发送验证码等等
func Register(c *gin.Context) {
	u := model.RegisterUser{}
	u.UserName = c.PostForm("userName")
	u.Password = c.PostForm("password")
	u.Email = c.PostForm("email")
	if u.UserName == "" {
		tool.RespErrWithData(c, false, "用户名不能为空")
		return
	}
	if len(u.UserName) >= 20 {
		tool.RespErrWithData(c, false, "用户名太长了")
		return
	}
	if len(u.Password) <= 6 {
		tool.RespErrWithData(c, false, "密码不能小于6个字符")
		return
	}
	if len(u.Password) >= 16 {
		tool.RespErrWithData(c, false, "密码不能大于16个字符")
		return
	}
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
	u.VerifyCode = c.PostForm("verifyCode")

	if u.VerifyCode == "" {
		tool.RespErrWithData(c, false, "请输入验证码")
		return
	}
	IsCorrect, err := service.CheckVerifyCodeByEmail(u)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "未发送验证码")
		return
	}
	if !IsCorrect {
		tool.RespErrWithData(c, false, "验证码错误")
		return
	}
	tool.RespErrWithData(c, true, "注册成功！")
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
