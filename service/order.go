package service

import (
	"JD/dao"
	"JD/model"
	"JD/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateOrder(o model.Order) (string, bool, error) {
	for _, v := range o.Settlement {
		//查询价格
		_, _, iG, err := dao.GetGoodsBaseInfo(v.GId)
		if err != nil {
			fmt.Println(err)
		}
		//赋值价格
		v.Price = iG.Price
		v.Name = iG.Name
		o.TotalPrice += v.Price * v.Account
	}

	//创建订单
	o.OrderNumber = tool.CreateOrder(time.Now())

	flag, err := dao.CreateOrder(o)
	return o.OrderNumber, flag, err
}

func CheckAllOrder(o model.Order) (map[int]model.Order, error) {
	return dao.CheckAllOrder(o)
}

func CheckSpecified(order model.Order) (model.Order, error) {
	return dao.CheckSpecified(order)
}

func CancelOrder(order model.Order) error {
	return dao.CancelOrder(order)
}

func PayOrder(c *gin.Context, order model.Order) (bool, error) {
	o, err := dao.CheckSpecified(order)
	if err != nil {
		return false, err
	}
	u := model.User{Id: order.Uid}
	u, err = dao.SearchUserByUid(u)
	if err != nil {
		return false, err
	}
	if u.Money < o.TotalPrice {
		return false, nil
	}
	//将所有settlement里的商品的销售量+account,扣钱,并将订单的状态进行修改
	flag, err := dao.SolveOrder(o, u)
	if err != nil {
		return false, err
	}
	if flag == false {
		tool.RespErrWithData(c, false, "订单错误!")
		c.Abort()
	}
	return true, nil
}

func CreateConsigneeInfo(consignee model.ConsigneeInfo) error {
	return dao.CreateConsigneeInfo(consignee)
}

func GetConsigneeInfo(consignee model.ConsigneeInfo) (map[int]model.ConsigneeInfo, error) {
	return dao.GetConsigneeInfo(consignee)
}

func DeleteConsigneeInfo(consignee model.ConsigneeInfo) error {
	return dao.DeleteConsigneeInfo(consignee)
}

func ConfirmOrder(order model.Order) error {
	return dao.ConfirmOrder(order)
}

func DeleteOrder(order model.Order) error {
	return dao.DeleteOrder(order)
}

func CheckOrderByStatus(order model.Order, status string) (map[int]model.Order, error) {
	return dao.CheckOrderByStatus(order, status)
}
