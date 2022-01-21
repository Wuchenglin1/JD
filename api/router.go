package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	Engine := gin.Default()

	VerifyCode := Engine.Group("/verify")
	{
		VerifyCode.POST("/sms/register", RegisterSendSMS)
		VerifyCode.POST("/email/register", RegisterSendEmail)
	}

	Check := Engine.Group("/check")
	{
		Check.POST("/sms/register", CheckRegisterSMS)
	}

	User := Engine.Group("/user")
	{
		User.POST("/login/normal", Login)
		User.POST("/register/email", Register)
	}
	_ = Engine.Run()
}
