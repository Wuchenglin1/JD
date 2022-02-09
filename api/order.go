package api

import (
	"JD/model"
	"JD/service"
	"JD/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateOrder(c *gin.Context) {
	//解析token
	claim, err := service.ParseAccessToken(c.PostForm("token"))
	if err != nil {
		if err.Error() == "token contains an invalid number of segments" {
			c.JSON(200, "token错误！")
			return
		}
	}
	flag, str := service.CheckTokenErr(claim, err)
	if !flag {
		tool.RespErrWithData(c, false, str)
		return
	}
	order := model.Order{
		Uid:        claim.User.Id,
		TotalPrice: 0,
		Status:     "等待付款",
		Time:       time.Now(),
	}

	address := c.PostForm("address")
	if address == "" {
		tool.RespErrWithData(c, false, "地址不能为空")
		return
	}
	order.Address = address

	phone := c.PostForm("phone")
	if phone == "" {
		tool.RespErrWithData(c, false, "电话不能为空")
		return
	}
	order.Phone = phone

	name := c.PostForm("name")
	if name == "" {
		tool.RespErrWithData(c, false, "联系人名字不能为空")
		return
	}
	order.Consignee = name

	payWay := c.PostForm("payWay")
	if payWay == "" {
		tool.RespErrWithData(c, false, "支付方式不正确")
		return
	}
	order.PayWay = payWay

	if err = c.ShouldBind(&order.Settlement); err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "商品格式有误")
		return
	}

	orderNumber, err := service.CreateOrder(order)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}

	tool.RespSuccessWithData(c, orderNumber)
}

func CheckAllOrder(c *gin.Context) {
	//解析token
	claim, err := service.ParseAccessToken(c.PostForm("token"))
	if err != nil {
		if err.Error() == "token contains an invalid number of segments" {
			c.JSON(200, "token错误！")
			return
		}
	}
	flag, str := service.CheckTokenErr(claim, err)
	if !flag {
		tool.RespErrWithData(c, false, str)
		return
	}
	order := model.Order{
		Uid: claim.User.Id,
	}

	m, err := service.CheckAllOrder(order)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "您的订单空空如也")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	for _, v := range m {
		c.JSON(200, v)
	}
}

func CheckSpecified(c *gin.Context) {
	//解析token
	claim, err := service.ParseAccessToken(c.PostForm("token"))
	if err != nil {
		if err.Error() == "token contains an invalid number of segments" {
			c.JSON(200, "token错误！")
			return
		}
	}
	flag, str := service.CheckTokenErr(claim, err)
	if !flag {
		tool.RespErrWithData(c, false, str)
		return
	}
	order := model.Order{}
	order.OrderNumber = c.PostForm("order")
	if order.OrderNumber == "" {
		tool.RespErrWithData(c, false, "订单号不能为空")
		return
	}
	order, err = service.CheckSpecified(order)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "该订单不存在")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	c.JSON(200, order)
}

func CancelOrder(c *gin.Context) {
	//解析token
	claim, err := service.ParseAccessToken(c.PostForm("token"))
	if err != nil {
		if err.Error() == "token contains an invalid number of segments" {
			c.JSON(200, "token错误！")
			return
		}
	}
	flag, str := service.CheckTokenErr(claim, err)
	if !flag {
		tool.RespErrWithData(c, false, str)
		return
	}

	order := model.Order{
		OrderNumber: c.PostForm("order"),
	}
	if order.OrderNumber == "" {
		tool.RespErrWithData(c, false, "订单不能为空")
		return
	}
	err = service.CancelOrder(order)

	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "该订单不存在")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccess(c)
}

func PayOrder(c *gin.Context) {
	//解析token
	claim, err := service.ParseAccessToken(c.PostForm("token"))
	if err != nil {
		if err.Error() == "token contains an invalid number of segments" {
			c.JSON(200, "token错误！")
			return
		}
	}
	flag, str := service.CheckTokenErr(claim, err)
	if !flag {
		tool.RespErrWithData(c, false, str)
		return
	}

	order := model.Order{
		OrderNumber: c.PostForm("order"),
		Uid:         claim.User.Id,
	}
	if order.OrderNumber == "" {
		tool.RespErrWithData(c, false, "订单号不能为空")
		return
	}
	flag, err = service.PayOrder(order)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "该订单不存在")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	if !flag {
		tool.RespErrWithData(c, false, "您的余额不足,请充值")
		return
	}
	tool.RespSuccessWithData(c, "支付成功")
}
