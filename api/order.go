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

	orderNumber, flag, err := service.CreateOrder(order)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	if flag == false {
		tool.RespErrWithData(c, false, "商品库存不足")
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
	//这里是当时写的时候没注意，然后如果要改就是大改，太麻烦了！我就直接把上下文传进去了
	flag, err = service.PayOrder(c, order)
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

func CreateConsigneeInfo(c *gin.Context) {
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

	consignee := model.ConsigneeInfo{
		Uid:            claim.User.Id,
		Name:           c.PostForm("name"),
		Province:       c.PostForm("province"),
		City:           c.PostForm("city"),
		Area:           c.PostForm("area"),
		Town:           c.PostForm("town"),
		DetailAddress:  c.PostForm("detailAddress"),
		Phone:          c.PostForm("phone"),
		FixedTelephone: c.PostForm("fixedTelephone"),
		Email:          c.PostForm("email"),
		AddressAlias:   c.PostForm("addressAlias"),
	}
	if consignee.Name == "" || consignee.Province == "" || consignee.City == "" || consignee.Area == "" || consignee.Town == "" || consignee.DetailAddress == "" || consignee.Phone == "" {
		tool.RespErrWithData(c, false, "参数不能为空")
		return
	}

	err = service.CreateConsigneeInfo(consignee)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccessWithData(c, "添加成功")
}

func GetConsigneeInfo(c *gin.Context) {
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

	consignee := model.ConsigneeInfo{Uid: claim.User.Id}
	m, err := service.GetConsigneeInfo(consignee)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "您还没有添加过收件人信息")
			return
		}
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	for _, v := range m {
		c.JSON(200, gin.H{
			"uid":            v.Uid,
			"cid":            v.Cid,
			"name":           v.Name,
			"address":        v.Province + "" + v.City + "" + v.Area + "" + v.Town,
			"detailAddress":  v.DetailAddress,
			"phone":          v.Phone,
			"fixedTelephone": v.FixedTelephone,
			"email":          v.Email,
			"addressAlias":   v.AddressAlias,
		})
	}
}

func DeleteConsigneeInfo(c *gin.Context) {
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

	consignee := model.ConsigneeInfo{Uid: claim.User.Id}
	err = service.DeleteConsigneeInfo(consignee)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "该收货人不存在")
			return
		}
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccessWithData(c, "删除成功")
}

func ConfirmOrder(c *gin.Context) {
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
		Uid:         claim.User.Id,
		OrderNumber: c.PostForm("order"),
	}

	err = service.ConfirmOrder(order)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "该订单不存在")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccessWithData(c, "取消成功")
}

func DeleteOrder(c *gin.Context) {
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
		Uid:         claim.User.Id,
		OrderNumber: c.PostForm("order"),
	}
	err = service.DeleteOrder(order)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "该订单不存在")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccessWithData(c, "删除成功")
}

func CheckOrderByStatus(c *gin.Context) {
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

	status := c.PostForm("status")
	switch status {
	case "1":
		status = "待付款"
	case "2":
		status = "待收货"
	case "3":
		status = "已完成"
	case "4":
		status = "已取消"
	default:
		tool.RespErrWithData(c, false, "status错误")
		return
	}
	order := model.Order{
		Uid: claim.User.Id,
	}
	m, err := service.CheckOrderByStatus(order, status)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "该订单不存在")
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
