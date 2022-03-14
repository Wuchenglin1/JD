package api

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {

	Engine := gin.Default()

	//Engine.Use(TlsHandler())

	Engine.Use(static.Serve("/", static.LocalFile("../static", false)))

	//开启中间件，允许跨域
	Engine.Use(Cors())

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
		User.POST("/rechargeBalance", RechargeBalance)
		User.GET("/checkBalance", CheckBalance)
		User.GET("/shoppingCart", BrowseShoppingCart)
	}

	Goods := Engine.Group("/goods")
	{
		Goods.POST("/create", Create)
		Goods.POST("/createGoods", CreateGoods)
		Goods.POST("/create/size", Size)
		Goods.POST("/photo/color", ColorPhoto)
		Goods.POST("/blouse", Blouse)
		Goods.POST("/add/shoppingCart", AddShoppingCart)
		Goods.POST("/browse/all", BrowseGoodsByKeyWords)
		Goods.POST("/focus", GoodsFocus)
		Goods.GET("/browse", BrowseGoods)
		Goods.GET("/getInfo", GetGoodsBaseInfo)
		Goods.GET("/getSize", GetGoodsSize)
		Goods.GET("/getColor", GetGoodsColor)
		Goods.GET("/browse/type", BrowseGoodsType)
		Goods.GET("/getFocus", GetGoodsFocus)
		Goods.DELETE("/delete/shoppingCart", DeleteShoppingCart)
		Goods.DELETE("/delete/focus", DeleteFocus)
	}

	order := Engine.Group("/order")
	{
		order.POST("/create", CreateOrder)
		order.POST("/createConsigneeInfo", CreateConsigneeInfo)
		order.PUT("/cancel", CancelOrder)
		order.PUT("/confirm", ConfirmOrder)
		order.POST("/pay", PayOrder)
		order.GET("/GetConsigneeInfo", GetConsigneeInfo)
		order.GET("/checkAll", CheckAllOrder)
		order.GET("/checkByStatus", CheckOrderByStatus)
		order.GET("/checkSpecified", CheckSpecified)
		order.DELETE("/delete", DeleteOrder)
		order.DELETE("/DeleteConsigneeInfo", DeleteConsigneeInfo)

	}

	comment := Engine.Group("/comment")
	{
		comment.POST("/add", AddComment)
		comment.POST("/reply", ReplyComment)
		comment.GET("/viewComment", ViewComment)
		comment.GET("/viewComment/specific", ViewSpecificComment)
	}

	store := Engine.Group("/store")
	{
		store.GET("/getGoods", GetGoods)
		store.GET("/getAnnouncement", GetAnnouncement)
		store.PUT("/postAnnouncement", UpdateAnnouncement)
	}
	token := Engine.Group("/token")
	{
		token.POST("/get", GetToken)
	}

	oauth := Engine.Group("/oauth")
	{
		oauth.GET("/", Authorize)
		oauth.GET("/authorization", Authorization)
	}

	_ = Engine.Run()
	// = Engine.RunTLS("", "/root/GoProject/csr/wuchenglin.plus_bundle.pem", "/root/GoProject/csr/wuchenglin.plus.key")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

//func TlsHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		secureMiddleware := secure.New(secure.Options{
//			SSLRedirect: true,
//			SSLHost:     "localhost:8080",
//		})
//		err := secureMiddleware.Process(c.Writer, c.Request)
//
//		// If there was an error, do not continue.
//		if err != nil {
//			return
//		}
//
//		c.Next()
//	}
//}
