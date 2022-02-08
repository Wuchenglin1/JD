package service

import (
	"JD/dao"
	"JD/model"
	"JD/tool"
	"fmt"
	"time"
)

func CreateOrder(o model.Order) (string, error) {
	for _, v := range o.Settlement {
		//查询价格
		iG, err := dao.GetGoodsBaseInfo(v.GId)
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

	err := dao.CreateOrder(o)
	return o.OrderNumber, err
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
