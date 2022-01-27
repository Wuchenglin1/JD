package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	Engine := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8080"}
	Engine.Use(cors.New(config))
	Engine.Static("./static", "./static")
	Engine.Use(static.Serve("/", static.LocalFile("../static", false)))

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

	Goods := Engine.Group("/goods")
	{
		Goods.POST("/create", Create)
		Goods.POST("/create/size", Size)
		Goods.POST("/photo/color", ColorPhoto)
		Goods.POST("/blouse", Blouse)
		Goods.GET("/browse", BrowseGoods)
		Goods.GET("/getInfo", GetGoodsBaseInfo)
		Goods.GET("/getSize", GetGoodsSize)
		Goods.GET("/getColor", GetGoodsColor)
	}

	token := Engine.Group("/token")
	{
		token.POST("/get", GetToken)
	}
	_ = Engine.Run()
}
